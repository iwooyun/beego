package adapter

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

type ApolloLogger struct{}

func (ApolloLogger) Debugf(format string, params ...interface{}) {
	logs.Debug(format, params)
}

func (ApolloLogger) Infof(format string, params ...interface{}) {
	logs.Info(format, params)
}

func (ApolloLogger) Warnf(format string, params ...interface{}) error {
	logs.Warn(format, params)
	return nil
}

func (ApolloLogger) Errorf(format string, params ...interface{}) error {
	logs.Error(format, params)
	return fmt.Errorf(format, params)
}

func (ApolloLogger) Debug(v ...interface{}) {
	logs.Debug(v)
}

func (ApolloLogger) Info(v ...interface{}) {
	logs.Info(v)
}

func (ApolloLogger) Warn(v ...interface{}) error {
	logs.Warn(v)
	return nil
}

func (ApolloLogger) Error(v ...interface{}) error {
	logs.Error(v)
	return fmt.Errorf("%s", v)
}
