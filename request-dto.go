package erpc

type RequestDto struct {
	Data interface{} `json:"data"`
	Hash string      `json:"hash"`
}
