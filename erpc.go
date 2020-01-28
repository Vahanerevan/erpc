package erpc

const StatusOK = "OK"
const StatusFAIL = "FAIL"

func IsStatusOK(status string) bool {
	return status == StatusOK
}
