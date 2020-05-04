package erpc

import (
	"encoding/json"
	"errors"
	"fmt"
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
	req.Req
	config Config
}

func (request *Request) SetPayload(dataObject interface{}) {

	request.payload = dataObject
}

func (request *Request) Call(action string, requestObject interface{}, path ...string) (*Response, error) {

	request.SetPayload(requestObject)

	bytes, err := ToJsonBytes(request.payload)

	if nil != err {
		return nil, err
	}
	header := req.Header{
		"Content-Type": "application/json",
		AuthHeader:     HashCalculate(bytes, request.config.Secret),
	}

	pathList := []string{request.config.URL, action}

	pathList = append(pathList, path...)

	uri := strings.Join(pathList, "/")

	if request.config.Debug {
		fmt.Println(uri, header, string(bytes))
	}

	resp, err := request.Post(uri, header, bytes)

	if nil != err {
		return nil, err
	}

	httpResp := resp.Response()

	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.New("response status is not 200")
	}

	responseData := &Response{Resp: resp}

	err = resp.ToJSON(responseData)

	if nil != err {
		return nil, err
	}

	if responseData.IsFail() {

		switch responseData.Code {
		case ErrorCodeHash:
			return nil, ErrInvalidHash
		}
		return nil, errors.New("Status undefined"+responseData.Message)
	}
	return responseData, nil
}

func (request *Request) validate() {}
