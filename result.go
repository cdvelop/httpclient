package httpclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func (h HttpClient) resultOK(p []js.Value) error {

	h.Log(p[0])

	msg := p[0].Get("statusText").String() //Not Found

	status := p[0].Get("status").String() //<number: 404>
	if status == "<number: 404>" {
		msg += " 404"
	}

	ok := p[0].Get("ok").String() //<boolean: false>
	if ok == "<boolean: false)>" {
		ok = "false"
	}

	// h.Log("RESP OK:", ok, "status:", status, "text", text)

	if len(p) != 2 {
		return model.Error(msg)
	}

	return nil
}
