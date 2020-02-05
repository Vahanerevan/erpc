package erpc

import (
	"encoding/json"
	"errors"
)

func NewParser() *Parser {
	return new(Parser)
}

type Parser struct {
	Config Config
	data   interface{}
	hash   string
}

func (parser *Parser) Configure(config Config) {
	parser.Config = config
}

func (parser *Parser) Handle(bytes []byte) error {
	var request RequestDto
	err := json.Unmarshal(bytes, &request)

	if nil != err {
		return err
	}

	parser.hash = request.Auth
	parser.data = request.Data

	return nil
}

func (parser *Parser) ToString() (string, error) {
	bytes, err := parser.ToBytes()
	if nil != err {
		return "", err
	}
	str := string(bytes)
	return str, nil
}

func (parser *Parser) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(parser.data)
	if nil != err {
		return nil, err
	}
	return bytes, nil
}

func (parser *Parser) ToJSON(object interface{}) error {
	bytes, err := parser.ToBytes()
	if nil != err {
		return err
	}
	return json.Unmarshal(bytes, object)
}

func (parser *Parser) Validate() error {
	dataString, err := parser.ToString()
	if nil != err {
		return err
	}
	localHash := HashCalculate(dataString, parser.Config.Secret)
	if localHash != parser.hash {
		return errors.New("Hash validation failed")
	}
	return nil
}
