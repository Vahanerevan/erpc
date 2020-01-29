package erpc

import (
	"encoding/json"
	"errors"
)

func NewReceiver() *Receiver {
	r := new(Receiver)
	return r
}

type Receiver struct {
	Config Config
	data   interface{}
	hash   string
}

func (receiver *Receiver) Configure(config Config) {
	receiver.Config = config
}

func (receiver *Receiver) Handle(bytes []byte) error {
	var request RequestDto
	err := json.Unmarshal(bytes, &request)

	if nil != err {
		return err
	}
	if nil != err {
		return err
	}

	receiver.hash = request.Hash
	receiver.data = request.Data

	return nil
}

func (receiver *Receiver) ToString() (string, error) {
	bytes, err := receiver.ToBytes()
	if nil != err {
		return "", err
	}
	str := string(bytes)
	return str, nil
}

func (receiver *Receiver) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(receiver.data)
	if nil != err {
		return nil, err
	}
	return bytes, nil
}

func (receiver *Receiver) ToJSON(object interface{}) error {
	bytes, err := receiver.ToBytes()
	if nil != err {
		return err
	}
	return json.Unmarshal(bytes, object)
}

func (receiver *Receiver) Validate() error {
	dataString, err := receiver.ToString()
	if nil != err {
		return err
	}
	localHash := HashCalculate(dataString, receiver.Config.Secret)
	if localHash != receiver.hash {
		return errors.New("Hash validation failed")
	}
	return nil
}
