package httpclient

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

type HttpClient struct {
	model.Logger
	*cutkey.Cut
	onComplete js.Func // Declarar la función onComplete fuera de fetchData
}
