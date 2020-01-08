package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

// CreateAlarmSign 告警平台加密规则
func CreateAlarmSign(param map[string]string, privateKey string) (token string) {
	keys := make([]string, 0, len(param))
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str string
	for _, k := range keys {
		str1 := k + "=" + param[k] + "&"
		str += str1
	}

	h := md5.New()
	h.Write([]byte(str + privateKey))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)

}
