package erpc

import "errors"

const StatusOK = "OK"
const StatusFAIL = "FAIL"
const XHeader string = "X-Auth"

const (
	ErrorCodeHash    int = 501
	ErrorCodeGeneral     = 505
)

var ErrInvalidHash = errors.New("Invalid Hash")

func IsStatusOK(status string) bool {
	return status == StatusOK
}

func ValidateInput(secret, hash string, input []byte) bool {
	return hash == HashCalculate(input, secret)
}
