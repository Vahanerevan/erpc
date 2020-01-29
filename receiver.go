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
	data   []byte
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
	receiver.data = bytes

	return nil
}

func (receiver *Receiver) ToString() string {
	return string(receiver.data)
}

func (receiver *Receiver) ToJSON(object interface{}) error {
	return json.Unmarshal(receiver.data, object)
}

func (receiver *Receiver) Validate() error {
	localHash := HashCalculate(string(receiver.data), receiver.Config.Secret)
	if localHash != receiver.hash {
		return errors.New("Hash validation failed")
	}
	return nil
}
