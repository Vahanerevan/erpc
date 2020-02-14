package erpc

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func HashCalculate(input []byte, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write(input)
	return hex.EncodeToString(h.Sum(nil))
}
