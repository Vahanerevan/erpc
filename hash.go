package erpc

import (
	"crypto/md5"
	"encoding/hex"
)

func HashCalculate(data string, secret string) string {
	text := []byte((data + secret))
	hashCalculator := md5.New()
	hashCalculator.Write([]byte(text))
	return hex.EncodeToString(hashCalculator.Sum(nil))
}
