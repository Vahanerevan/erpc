package erpc

import "errors"

type Status string

const StatusOK Status = "OK"
const StatusFAIL Status = "FAIL"

const AuthHeader string = "X-Auth"

const (
	ErrorCodeHash    int = 501
	ErrorCodeGeneral     = 505
)

var ErrInvalidHash = errors.New("Invalid Hash")

func IsStatusOK(status Status) bool {
	return status == StatusOK
}

func ValidateInput(secret, hash string, input []byte) bool {
	return hash == HashCalculate(input, secret)
}
