/**
 * @Author: lixiang
 * @Date: 2025-03-23 17:02
 * @Description: md5.go
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	cipher := h.Sum(nil)
	return hex.EncodeToString(cipher)
}
