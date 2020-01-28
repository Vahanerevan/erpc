package erpc

import (
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strings"
)

type RequestDto struct {
	Data interface{} `json:"data"`
	Hash string      `json:"hash"`
}

func NewRequest(config Config) *Request {
	return &Request{config: config}
}

type Request struct {
	requestData RequestDto
	Resp        Response
	config      Config
}

func (request *Request) SetRequestObject(dataObject interface{}) error {
	bytes, err := json.Marshal(dataObject)
	if nil != err {
		return err
	}
	hash := HashCalculate(string(bytes), request.config.Secret)

	request.requestData = RequestDto{
		Data: dataObject,
		Hash: hash,
	}
	return err
}

func (request *Request) Call(path ...string) (*Response, error) {

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
	return responseData, nil
}



func (request *Request) validate() {}
