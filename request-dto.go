package erpc

type RequestDto struct {
	Data interface{} `json:"data"`
	Auth string      `json:"auth"`
}
