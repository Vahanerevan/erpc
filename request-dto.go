package erpc

type RequestDto struct {
	Data   interface{} `json:"data"`
	Action string      `json:"action"`
}
