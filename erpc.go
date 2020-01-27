package erpc

const HashParameter string = "hash"
const StatusOK = "OK"
const StatusFAIL = "FAIL"

func IsStatusOK(status string) bool {
	return status == StatusOK
}
