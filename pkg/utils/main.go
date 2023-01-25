package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Contains(lst []string, val string) bool {
	for _, v := range lst {
		if v == val {
			return true
		}
	}
	return false
}

func SHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
