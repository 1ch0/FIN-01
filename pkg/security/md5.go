package security

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	_, _ = m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
