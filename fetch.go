package httpclient

import (
	"syscall/js"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

func (h *HttpClient) SendJson(o *model.Object, data []map[string]string, action string, out_resp func([]model.Response, error)) {

	body, err := cutkey.Encode(o, data...)
	if err != nil {
		out_resp(nil, err)
		return
	}

	endpoint := action + "/" + o.Name

	// h.Log("API endpoint:", endpoint)

	// Crear una función JavaScript que se llamará cuando se complete la solicitud
	h.onComplete = js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		// argumento 0 es el cuerpo de la respuesta de la solicitud Fetch, que debería ser una cadena de texto JSON.
		// argumento 1 indica si la promesa se resolvió o se rechazó.
		err = h.resultOK(p)
		if err != nil {
			out_resp(nil, err)
			return nil
		}
		h.Log("RESPUESTA:", p[0].String())

		// Decodificar la respuesta
		responseData := h.DecodeResponses([]byte(p[0].String()))

		// Llamar a la función de respuesta de Go con los datos decodificados
		out_resp(responseData, nil)

		// Liberar la función JavaScript
		h.onComplete.Release()

		return nil
	})

	// Realizar la solicitud Fetch en JavaScript
	js.Global().Get("fetch").Invoke(endpoint, js.ValueOf(map[string]interface{}{
		"method":  "POST",
		"body":    js.ValueOf(string(body)),
		"headers": js.ValueOf(map[string]interface{}{"Content-Type": "application/json"}),
	})).Call("then", h.onComplete, js.Null())

}
