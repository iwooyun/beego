// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const MultiFileAdapter = "AdapterMultiFile"

const (
	y1  = `0123456789`
	y2  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	y3  = `0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999`
	y4  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	mo1 = `000000000111`
	mo2 = `123456789012`
	d1  = `0000000001111111111222222222233`
	d2  = `1234567890123456789012345678901`
	h1  = `000000000011111111112222`
	h2  = `012345678901234567890123`
	mi1 = `000000000011111111112222222222333333333344444444445555555555`
	mi2 = `012345678901234567890123456789012345678901234567890123456789`
	s1  = `000000000011111111112222222222333333333344444444445555555555`
	s2  = `012345678901234567890123456789012345678901234567890123456789`
	ns1 = `0123456789`
)

func formatTimeHeader(when time.Time) ([]byte, int, int) {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	ns := when.Nanosecond() / 1000000
	var buf [24]byte

	buf[0] = y1[y/1000%10]
	buf[1] = y2[y/100]
	buf[2] = y3[y-y/100*100]
	buf[3] = y4[y-y/100*100]
	buf[4] = '/'
	buf[5] = mo1[mo-1]
	buf[6] = mo2[mo-1]
	buf[7] = '/'
	buf[8] = d1[d-1]
	buf[9] = d2[d-1]
	buf[10] = ' '
	buf[11] = h1[h]
	buf[12] = h2[h]
	buf[13] = ':'
	buf[14] = mi1[mi]
	buf[15] = mi2[mi]
	buf[16] = ':'
	buf[17] = s1[s]
	buf[18] = s2[s]
	buf[19] = '.'
	buf[20] = ns1[ns/100]
	buf[21] = ns1[ns%100/10]
	buf[22] = ns1[ns%10]

	buf[23] = ' '

	return buf[0:], d, h
}

// A filesLogWriter manages several fileLogWriter
// filesLogWriter will write logs to the file in json configuration  and write the same level log to correspond file
// means if the file name in configuration is project.log filesLogWriter will create project.error.log/project.debug.log
// and write the error-level logs to project.error.log and write the debug-level logs to project.debug.log
// the rotate attribute also  acts like fileLogWriter
type multiFileLogWriter struct {
	writers       [logs.LevelDebug + 1]*fileLogWriter // the last one for fullLogWriter
	fullLogWriter *fileLogWriter
	Separate      []string `json:"separate"`
}

var levelNames = [...]string{"access", "alert", "critical", "error", "warning", "notice", "info", "debug"}

// Init file logger with json config.
// jsonConfig like:
// 	{
// 	"filename":"logs/beego.log",
// 	"maxLines":0,
// 	"maxsize":0,
// 	"daily":true,
// 	"maxDays":15,
// 	"rotate":true,
//  	"perm":0600,
// 	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
// 	}
func (f *multiFileLogWriter) Init(config string) error {
	writer := newFileWriter().(*fileLogWriter)
	err := writer.Init(config)
	if err != nil {
		return err
	}
	f.fullLogWriter = writer

	// unmarshal "separate" field to f.Separate
	_ = json.Unmarshal([]byte(config), f)

	jsonMap := map[string]interface{}{}
	_ = json.Unmarshal([]byte(config), &jsonMap)

	for i := logs.LevelEmergency; i < logs.LevelDebug+1; i++ {
		for _, v := range f.Separate {
			if v == levelNames[i] {
				jsonMap["filename"] = f.fullLogWriter.fileNameOnly + "." + levelNames[i] + f.fullLogWriter.suffix
				jsonMap["level"] = i
				bs, _ := json.Marshal(jsonMap)
				writer = newFileWriter().(*fileLogWriter)
				err := writer.Init(string(bs))
				if err != nil {
					return err
				}
				f.writers[i] = writer
			}
		}
	}

	return nil
}

func (f *multiFileLogWriter) Destroy() {
	for i := 0; i < len(f.writers); i++ {
		if f.writers[i] != nil {
			f.writers[i].Destroy()
		}
	}
}

func (f *multiFileLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	needWrite := true
	for i := 0; i < len(f.writers)-1; i++ {
		if f.writers[i] != nil {
			if level == f.writers[i].Level {
				_ = f.writers[i].WriteMsg(when, msg, level)
				needWrite = false
			}
		}
	}
	if needWrite {
		_ = f.fullLogWriter.WriteMsg(when, msg, level)
	}

	return nil
}

func (f *multiFileLogWriter) Flush() {
	for i := 0; i < len(f.writers); i++ {
		if f.writers[i] != nil {
			f.writers[i].Flush()
		}
	}
}

// newFilesWriter create a FileLogWriter returning as LoggerInterface.
func newFilesWriter() logs.Logger {
	return &multiFileLogWriter{}
}

type fileLogWriter struct {
	sync.RWMutex // write log order by order and  atomic incr maxLinesCurLines and maxSizeCurSize
	// The opened file
	Filename   string `json:"filename"`
	fileWriter *os.File

	// Rotate at line
	MaxLines         int `json:"maxlines"`
	maxLinesCurLines int

	MaxFiles         int `json:"maxfiles"`
	MaxFilesCurFiles int

	// Rotate at size
	MaxSize        int `json:"maxsize"`
	maxSizeCurSize int

	// Rotate daily
	Daily         bool  `json:"daily"`
	MaxDays       int64 `json:"maxdays"`
	dailyOpenDate int
	dailyOpenTime time.Time

	// Rotate hourly
	Hourly         bool  `json:"hourly"`
	MaxHours       int64 `json:"maxhours"`
	hourlyOpenDate int
	hourlyOpenTime time.Time

	Rotate bool `json:"rotate"`

	Level int `json:"level"`

	Perm string `json:"perm"`

	RotatePerm string `json:"rotateperm"`

	fileNameOnly, suffix string // like "project.log", project is fileNameOnly and .log is suffix
}

// newFileWriter create a FileLogWriter returning as LoggerInterface.
func newFileWriter() logs.Logger {
	w := &fileLogWriter{
		Daily:      true,
		MaxDays:    7,
		Hourly:     false,
		MaxHours:   168,
		Rotate:     true,
		RotatePerm: "0440",
		Level:      logs.LevelTrace,
		Perm:       "0660",
		MaxLines:   10000000,
		MaxFiles:   999,
		MaxSize:    1 << 28,
	}
	return w
}

// Init file logger with json config.
// jsonConfig like:
//  {
//  "filename":"logs/beego.log",
//  "maxLines":10000,
//  "maxsize":1024,
//  "daily":true,
//  "maxDays":15,
//  "rotate":true,
//      "perm":"0600"
//  }
func (w *fileLogWriter) Init(jsonConfig string) error {
	err := json.Unmarshal([]byte(jsonConfig), w)
	if err != nil {
		return err
	}
	if len(w.Filename) == 0 {
		return errors.New("jsonconfig must have filename")
	}
	w.suffix = filepath.Ext(w.Filename)
	w.fileNameOnly = strings.TrimSuffix(w.Filename, w.suffix)
	if w.suffix == "" {
		w.suffix = ".log"
	}
	err = w.startLogger()
	return err
}

// start file logger. create log file and set to locker-inside file writer.
func (w *fileLogWriter) startLogger() error {
	file, err := w.createLogFile()
	if err != nil {
		return err
	}
	if w.fileWriter != nil {
		w.fileWriter.Close()
	}
	w.fileWriter = file
	return w.initFd()
}

func (w *fileLogWriter) needRotateDaily(size int, day int) bool {
	return (w.MaxLines > 0 && w.maxLinesCurLines >= w.MaxLines) ||
		(w.MaxSize > 0 && w.maxSizeCurSize >= w.MaxSize) ||
		(w.Daily && day != w.dailyOpenDate)
}

func (w *fileLogWriter) needRotateHourly(size int, hour int) bool {
	return (w.MaxLines > 0 && w.maxLinesCurLines >= w.MaxLines) ||
		(w.MaxSize > 0 && w.maxSizeCurSize >= w.MaxSize) ||
		(w.Hourly && hour != w.hourlyOpenDate)

}

// WriteMsg write logger message into file.
func (w *fileLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > w.Level {
		return nil
	}

	hd, d, h := formatTimeHeader(when)
	if level != logs.LevelEmergency {
		msg = string(hd) + msg + "\n"
	} else {
		msg = strings.TrimSpace(strings.TrimPrefix(msg, "[M]")) + "\n"
	}

	if w.Rotate {
		w.RLock()
		if w.needRotateHourly(len(msg), h) {
			w.RUnlock()
			w.Lock()
			if w.needRotateHourly(len(msg), h) {
				if err := w.doRotate(when); err != nil {
					logs.Error("FileLogWriter(%q): %s\n", w.Filename, err)
				}
			}
			w.Unlock()
		} else if w.needRotateDaily(len(msg), d) {
			w.RUnlock()
			w.Lock()
			if w.needRotateDaily(len(msg), d) {
				if err := w.doRotate(when); err != nil {
					logs.Error("FileLogWriter(%q): %s\n", w.Filename, err)
				}
			}
			w.Unlock()
		} else {
			w.RUnlock()
		}
	}

	w.Lock()
	_, err := w.fileWriter.Write([]byte(msg))
	if err == nil {
		w.maxLinesCurLines++
		w.maxSizeCurSize += len(msg)
	}
	w.Unlock()
	return err
}

func (w *fileLogWriter) createLogFile() (*os.File, error) {
	// Open the log file
	perm, err := strconv.ParseInt(w.Perm, 8, 64)
	if err != nil {
		return nil, err
	}

	filepathDir := path.Dir(w.Filename)
	_ = os.MkdirAll(filepathDir, os.FileMode(perm))

	fd, err := os.OpenFile(w.Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(perm))
	if err == nil {
		// Make sure file perm is user set perm cause of `os.OpenFile` will obey umask
		_ = os.Chmod(w.Filename, os.FileMode(perm))
	}
	return fd, err
}

func (w *fileLogWriter) initFd() error {
	fd := w.fileWriter
	fInfo, err := fd.Stat()
	if err != nil {
		return fmt.Errorf("get stat err: %s", err)
	}
	w.maxSizeCurSize = int(fInfo.Size())
	w.dailyOpenTime = time.Now()
	w.dailyOpenDate = w.dailyOpenTime.Day()
	w.hourlyOpenTime = time.Now()
	w.hourlyOpenDate = w.hourlyOpenTime.Hour()
	w.maxLinesCurLines = 0
	if w.Hourly {
		go w.hourlyRotate(w.hourlyOpenTime)
	} else if w.Daily {
		go w.dailyRotate(w.dailyOpenTime)
	}
	if fInfo.Size() > 0 && w.MaxLines > 0 {
		count, err := w.lines()
		if err != nil {
			return err
		}
		w.maxLinesCurLines = count
	}
	return nil
}

func (w *fileLogWriter) dailyRotate(openTime time.Time) {
	y, m, d := openTime.Add(24 * time.Hour).Date()
	nextDay := time.Date(y, m, d, 0, 0, 0, 0, openTime.Location())
	tm := time.NewTimer(time.Duration(nextDay.UnixNano() - openTime.UnixNano() + 100))
	<-tm.C
	w.Lock()
	if w.needRotateDaily(0, time.Now().Day()) {
		if err := w.doRotate(time.Now()); err != nil {
			logs.Error("FileLogWriter(%q): %s\n", w.Filename, err)
		}
	}
	w.Unlock()
}

func (w *fileLogWriter) hourlyRotate(openTime time.Time) {
	y, m, d := openTime.Add(1 * time.Hour).Date()
	h, _, _ := openTime.Add(1 * time.Hour).Clock()
	nextHour := time.Date(y, m, d, h, 0, 0, 0, openTime.Location())
	tm := time.NewTimer(time.Duration(nextHour.UnixNano() - openTime.UnixNano() + 100))
	<-tm.C
	w.Lock()
	if w.needRotateHourly(0, time.Now().Hour()) {
		if err := w.doRotate(time.Now()); err != nil {
			logs.Error("FileLogWriter(%q): %s\n", w.Filename, err)
		}
	}
	w.Unlock()
}

func (w *fileLogWriter) lines() (int, error) {
	fd, err := os.Open(w.Filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()

	buf := make([]byte, 32768) // 32k
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := fd.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}

		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			break
		}
	}

	return count, nil
}

// DoRotate means it need to write file in new file.
// new file name like xx.2013-01-01.log (daily) or xx.001.log (by line or size)
func (w *fileLogWriter) doRotate(logTime time.Time) error {
	// file exists
	// Find the next available number
	num := w.MaxFilesCurFiles + 1
	fName := ""
	format := ""
	var openTime time.Time
	rotatePerm, err := strconv.ParseInt(w.RotatePerm, 8, 64)
	if err != nil {
		return err
	}

	_, err = os.Lstat(w.Filename)
	if err != nil {
		// even if the file is not exist or other ,we should RESTART the logger
		goto RestartLogger
	}

	if w.Hourly {
		format = "2006010215"
		openTime = w.hourlyOpenTime
	} else if w.Daily {
		format = "2006-01-02"
		openTime = w.dailyOpenTime
	}

	// only when one of them be setted, then the file would be splited
	if w.MaxLines > 0 || w.MaxSize > 0 {
		for ; err == nil && num <= w.MaxFiles; num++ {
			fName = w.fileNameOnly + fmt.Sprintf(".%s.%03d%s", logTime.Format(format), num, w.suffix)
			_, err = os.Lstat(fName)
		}
	} else {
		fName = w.fileNameOnly + fmt.Sprintf(".%s.%03d%s", openTime.Format(format), num, w.suffix)
		_, err = os.Lstat(fName)
		w.MaxFilesCurFiles = num
	}

	// return error if the last file checked still existed
	if err == nil {
		return fmt.Errorf("Rotate: Cannot find free log number to rename %s ", w.Filename)
	}

	// close fileWriter before rename
	w.fileWriter.Close()

	// Rename the file to its new found name
	// even if occurs error,we MUST guarantee to  restart new logger
	err = os.Rename(w.Filename, fName)
	if err != nil {
		logs.Error("Rename: Cannot rename  %s  to %s, err => %s", w.Filename, fName, err.Error())
		goto RestartLogger
	}

	err = os.Chmod(fName, os.FileMode(rotatePerm))

RestartLogger:

	startLoggerErr := w.startLogger()
	go w.deleteOldLog()

	if startLoggerErr != nil {
		return fmt.Errorf("Rotate StartLogger: %s ", startLoggerErr)
	}
	if err != nil {
		return fmt.Errorf("Rotate: %s ", err)
	}
	return nil
}

func (w *fileLogWriter) deleteOldLog() {
	dir := filepath.Dir(w.Filename)
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) (returnErr error) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("Unable to delete old log '%s', error: %v\n", path, r)
			}
		}()

		if info == nil {
			return
		}
		if w.Hourly {
			if !info.IsDir() && info.ModTime().Add(1*time.Hour*time.Duration(w.MaxHours)).Before(time.Now()) {
				if strings.HasPrefix(filepath.Base(path), filepath.Base(w.fileNameOnly)) &&
					strings.HasSuffix(filepath.Base(path), w.suffix) {
					_ = os.Remove(path)
				}
			}
		} else if w.Daily {
			if !info.IsDir() && info.ModTime().Add(24*time.Hour*time.Duration(w.MaxDays)).Before(time.Now()) {
				if strings.HasPrefix(filepath.Base(path), filepath.Base(w.fileNameOnly)) &&
					strings.HasSuffix(filepath.Base(path), w.suffix) {
					_ = os.Remove(path)
				}
			}
		}
		return
	})
}

// Destroy close the file description, close file writer.
func (w *fileLogWriter) Destroy() {
	w.fileWriter.Close()
}

// Flush flush file logger.
// there are no buffering messages in file logger in memory.
// flush file means sync file from disk.
func (w *fileLogWriter) Flush() {
	_ = w.fileWriter.Sync()
}

func init() {
	logs.Register(MultiFileAdapter, newFilesWriter)
}
