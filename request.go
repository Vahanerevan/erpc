package erpc

import (
	//"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strings"
)

func NewRequest(config Config) *Request {
	return &Request{config: config}
}

type Request struct {
	requestData RequestDto
	Resp        Response
	config      Config
}

func (request *Request) SetRequestObject(dataObject interface{}) {
	//bytes, err := json.Marshal(dataObject)
	//if nil != err {
	//	return err
	//}
	//hash := HashCalculate(string(bytes), request.config.Secret)

	request.requestData = RequestDto{
		Data: dataObject,
		Auth: request.config.Secret,
	}
	//return err
}

func (request *Request) Call(requestObject interface{}, path ...string) (*Response, error) {
	request.SetRequestObject(requestObject)
	header := req.Header{
		"Content-Type": "application/json",
	}
	pathList := []string{request.config.URL}
	pathList = append(pathList, path...)
	uri := strings.Join(pathList, "/")
	r, err := req.Post(uri, header, req.BodyJSON(&request.requestData))

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
