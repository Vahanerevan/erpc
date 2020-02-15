package erpc

const StatusOK = "OK"
const StatusFAIL = "FAIL"
const XHeader string = "X-Auth"



func IsStatusOK(status string) bool {
	return status == StatusOK
}
