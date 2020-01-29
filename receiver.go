package erpc

import (
	"encoding/json"
	"errors"
)

func ReceiveBytes(bytes []byte) *Receiver {
	r := new(Receiver)
	r.Handle(bytes)
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
	dataString, err := json.Marshal(receiver.data)
	if nil != err {
		return nil, err
	}
	return dataString, nil
}

func (receiver *Receiver) ToJSON(object interface{}) error {
	str, err := receiver.ToBytes()
	if nil != err {
		return err
	}
	return json.Unmarshal(str, object)
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
