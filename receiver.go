package erpc

import (
	"encoding/json"
	"errors"
)

func NewReceiver(config Config) *Receiver {
	return &Receiver{Config: config}
}

type Receiver struct {
	Config Config
	data   []byte
}

func (receiver *Receiver) Handle(bytes []byte) error {
	var request RequestDto
	err := json.Unmarshal(bytes, &request)
	if nil != err {
		return err
	}
	dataString, err := json.Marshal(request.Data)
	if nil != err {
		return err
	}
	localHash := HashCalculate(string(dataString), receiver.Config.Secret)
	if localHash != request.Hash {
		return errors.New("Hash validation failed")
	}
	receiver.data = bytes

	return nil
}

func (receiver *Receiver) ToString() string {
	return string(receiver.data)
}

func (receiver *Receiver) ToJSON(object interface{}) {
	json.Unmarshal(receiver.data, object)
}
