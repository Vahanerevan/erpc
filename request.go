package erpc

import (
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strings"
)

func NewRequest(config RequestConfig) *Request {
	return &Request{config: config}
}

type Request struct {
	requestMap map[string]interface{}
	Resp       Response
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

func (request *Request) Call(path ...string) error {
	header := req.Header{
		"Content-Type": "application/json",
	}
	pathList := []string{request.config.URL}
	pathList = append(pathList, path...)
	uri := strings.Join(pathList, "/")
	r, err := req.Post(uri, header, req.BodyJSON(&request.requestMap))
	if nil != err {
		return err
	}
	resp := r.Response()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Response status is not 200")
	}

	erpcResp := &request.Resp
	r.ToJSON(erpcResp)
	erpcResp.Resp = r

	if false == request.IsOk() {
		return errors.New("Request failed on remote")
	}
	return nil
}

func (request *Request) IsOk() bool {
	return IsStatusOK(request.Resp.Status)
}

func (request *Request) validate() {}
