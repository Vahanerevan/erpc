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

const XAuthHash string = "X-AuthHash"

func NewRequest(config Config) *Request {
	return &Request{config: config}
}

func ToByteString(data interface{}) ([]byte, error) {
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
	}
}

func (request *Request) Action(action string, requestObject interface{}, path ...string) (*Response, error) {

	request.SetRequestObject(requestObject)
	request.payload.Action = action

	bytes, err := ToByteString(request.payload)

	header := req.Header{
		"Content-Type": "application/json",
		XAuthHash:      HashCalculate(bytes, request.config.Secret),
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
