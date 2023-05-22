package util

import (
	"crypto/md5"
	"encoding/hex"
)

//主要对文件名进行MD5加密
func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
