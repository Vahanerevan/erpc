package erpc

import (
	"encoding/json"
	"fmt"

	//"encoding/json"
	"errors"
	"github.com/imroc/req"
	"net/http"
	"strings"
)

func NewRequest(config Config) *Request {
	return &Request{config: config}
}

func ToJsonBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

type Request struct {
	payload interface{}
	Resp    Response
	config  Config
}

func (request *Request) SetPayload(dataObject interface{}) {

	request.payload = dataObject
}

func (request *Request) Call(action string, requestObject interface{}, path ...string) (*Response, error) {

	request.SetPayload(requestObject)

	bytes, err := ToJsonBytes(request.payload)

	header := req.Header{
		"Content-Type": "application/json",
		XHeader:        HashCalculate(bytes, request.config.Secret),
	}

	pathList := []string{request.config.URL, action}

	pathList = append(pathList, path...)

	uri := strings.Join(pathList, "/")

	fmt.Println(uri)

	r, err := req.Post(uri, header, bytes)

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
