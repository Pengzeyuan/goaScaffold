package util

import (
	"crypto/md5"
	"fmt"
)

func HashStrEncode(salt, raw string) string {
	// salt := config.C.Salt
	data := []byte(salt + raw)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
