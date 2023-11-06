package httpclient

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func Add(l model.Logger, c *cutkey.Cut) (*HttpClient, error) {

	n := HttpClient{
		Logger:     l,
		Cut:        c,
		onComplete: js.Func{},
	}

	return &n, nil
}
