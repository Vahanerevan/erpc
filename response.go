package erpc

import "github.com/imroc/req"

type Response struct {
	*req.Resp
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (response *Response) IsOk() bool {
	return IsStatusOK(response.Status)
}
func (response *Response) IsFail() bool {
	return false == IsStatusOK(response.Status)
}

func (response *Response) BindJSON(data interface{}) error {
	return response.ToJSON(data)
}
