package httpclient

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func Add(log model.Logger) (*HttpClient, error) {

	n := HttpClient{
		Logger:     nil,
		Cut:        &cutkey.Cut{},
		onComplete: js.Func{},
	}

	return &n, nil
}
