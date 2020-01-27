package erpc

import (
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
)

func NewRequest(config RequestConfig) *Request {
	return &Request{config: config}
}

type Request struct {
	requestMap map[string]interface{}
	response   Response
	config     RequestConfig
}

func (request *Request) SetRequestObject(object interface{}) error {
	bytes, err := json.Marshal(object)
	if nil != err {
		return err
	}
	hash := HashCalculate(string(bytes), request.config.Secret)
	mapData := make(map[string]interface{})
	json.Unmarshal(bytes, &mapData)
	mapData[HashParameter] = hash
	request.requestMap = mapData
	return err
}

func (request *Request) Call() error {
	header := req.Header{
		"Content-Type": "application/json",
	}
	r, err := req.Post(request.config.URL, header, req.BodyJSON(&request.requestMap))
	if nil != err {
		return err
	}
	resp := r.Response()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Response status is not 200")
	}

	erpcResp := &request.response
	r.ToJSON(erpcResp)
	erpcResp.Resp = r

	if false == request.IsOk() {
		return errors.New("Request failed on remote")
	}
	return nil
}

func (request *Request) Response() Response {
	return request.response
}

func (request *Request) IsOk() bool {
	return IsStatusOK(request.response.Status)
}

func (request *Request) validate() {}
