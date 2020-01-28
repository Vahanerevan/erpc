package erpc

import (
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strings"
)

type RequestDto struct {
	Data interface{}
	Hash string
}

func NewRequest(config RequestConfig) *Request {
	return &Request{config: config}
}

type Request struct {
	requestData RequestDto
	Resp        Response
	config      RequestConfig
}

func (request *Request) SetRequestObject(object interface{}) error {
	bytes, err := json.Marshal(object)
	if nil != err {
		return err
	}
	hash := HashCalculate(string(bytes), request.config.Secret)

	request.requestData = RequestDto{
		Data: object,
		Hash: hash,
	}
	return err
}

func (request *Request) Call(path ...string) error {

	header := req.Header{
		"Content-Type": "application/json",
	}
	pathList := []string{request.config.URL}
	pathList = append(pathList, path...)
	uri := strings.Join(pathList, "/")
	r, err := req.Post(uri, header, req.BodyJSON(&request.requestData))

	if nil != err {
		return err
	}

	resp := r.Response()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Response status is not 200")
	}

	responseData := &request.Resp
	r.ToJSON(responseData)
	responseData.Resp = r

	if false == request.IsOk() {
		return errors.New("Request failed on remote")
	}
	return nil
}

func (request *Request) IsOk() bool {
	return IsStatusOK(request.Resp.Status)
}

func (request *Request) validate() {}
