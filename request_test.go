package erpc

import (
	"github.com/imroc/req"
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Call(t *testing.T) {
	type fields struct {
		payload interface{}
		Req     req.Req
		config  Config
	}
	type args struct {
		action        string
		requestObject interface{}
		path          []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{{
		name: "test call",
		fields:fields{
			payload: struct {
				A string `json:"a"`
			}{A:"test"},
			Req:     req.Req{},
			config:  Config{
				Secret: "test",
				URL:    "http://localhost",
				Debug:  nil,
			},
		}, 
		// TODO: Add test cases.
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &Request{
				payload: tt.fields.payload,
				Req:     tt.fields.Req,
				config:  tt.fields.config,
			}
			got, err := request.Call(tt.args.action, tt.args.requestObject, tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Call() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetPayload(t *testing.T) {
	type fields struct {
		payload interface{}
		Req     req.Req
		config  Config
	}
	type args struct {
		dataObject interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Request{
				payload: tt.fields.payload,
				Req:     tt.fields.Req,
				config:  tt.fields.config,
			}
		})
	}
}

func TestRequest_validate(t *testing.T) {
	type fields struct {
		payload interface{}
		Req     req.Req
		config  Config
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Request{
				payload: tt.fields.payload,
				Req:     tt.fields.Req,
				config:  tt.fields.config,
			}
		})
	}
}

func TestToJsonBytes(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJsonBytes(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJsonBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJsonBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
