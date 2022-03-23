package cryptox

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
