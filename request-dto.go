package erpc

type RequestDTO struct {
	Data   interface{} `json:"data"`
	Action string      `json:"action"`
}
