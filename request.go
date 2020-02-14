package erpc

import (
	"encoding/json"
	//"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strings"
)

func NewRequest(config Config) *Request {
	return &Request{config: config}
}

func TyBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

type Request struct {
	payload RequestDto
	Resp    Response
	config  Config
}

func (request *Request) SetRequestObject(dataObject interface{}) {

	request.payload = RequestDto{
		Data: dataObject,
		Auth: request.config.Auth,
	}
}

func (request *Request) Call(requestObject interface{}, path ...string) (*Response, error) {
	request.SetRequestObject(requestObject)
	header := req.Header{
		"Content-Type": "application/json",
	}
	pathList := []string{request.config.URL}
	pathList = append(pathList, path...)
	uri := strings.Join(pathList, "/")
	r, err := req.Post(uri, header, req.BodyJSON(&request.payload))

	if nil != err {
		return nil, err
	}

	resp := r.Response()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Response status is not 200")
	}

	responseData := &request.Resp
	r.ToJSON(responseData)
	responseData.Resp = r

	return responseData, nil
}

func (request *Request) Action(action string, requestObject interface{}, path ...string) (*Response, error) {

	request.SetRequestObject(requestObject)


	if nil != err {
		return nil, err
	}

	request.payload.Action = action

	header := req.Header{
		"Content-Type": "application/json",
	}

	pathList := []string{request.config.URL}

	pathList = append(pathList, path...)

	uri := strings.Join(pathList, "/")

	r, err := req.Post(uri, header, req.BodyJSON(&requestObject))

	if nil != err {
		return nil, err
	}

	resp := r.Response()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Response status is not 200")
	}

	responseData := &request.Resp
	r.ToJSON(responseData)
	responseData.Resp = r

	return responseData, nil
}

func (request *Request) validate() {}
