package erpc

import "github.com/imroc/req"

type IStatus struct {
	Status string `json:"status"`
}

type Response struct {
	IStatus
	*req.Resp
}
