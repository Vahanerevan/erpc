package erpc

import (
	"encoding/json"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

type Response struct {
	*req.Resp
	Status  Status      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (response *Response) WithStatus(status Status) *Response {
	response.Status = status
	return response
}

func (response *Response) WithMessage(message string) *Response {
	response.Message = message
	return response
}
func (response *Response) WithCode(code int) *Response {
	response.Code = code
	return response
}

func (response *Response) WithData(data interface{}) *Response {
	response.Data = data
	return response
}

func (response *Response) ToJSON(in interface{}) error {
	result := gjson.Get(response.String(), "data")
	return json.Unmarshal([]byte(result.Raw), in)
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
