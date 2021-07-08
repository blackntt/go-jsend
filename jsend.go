package jsend

import (
	"encoding/json"
	"net/http"
)

//JSend ...
type JSend struct {
	data []byte
	code int
}

func (js *JSend) Send(w http.ResponseWriter) {
	w.WriteHeader(js.code)
	w.Write(js.data)
}

//IJSendBuilder ...
type IJSendBuilder interface {
	Code(int) IJSendBuilder
	Data(interface{}) IJSendBuilder
	Message(string) IJSendBuilder
	Build() *JSend
}

//JSendBuilder ...
type JSendBuilder struct {
	code    int
	data    interface{}
	message string
}

//Code ...
func (builder *JSendBuilder) Code(code int) IJSendBuilder {
	builder.code = code
	return builder
}

//Data ...
func (builder *JSendBuilder) Data(data interface{}) IJSendBuilder {
	builder.data = data
	return builder
}

//Message ...
func (builder *JSendBuilder) Message(message string) IJSendBuilder {
	builder.message = message
	return builder
}

//Build ...
func (builder *JSendBuilder) Build() *JSend {
	data := make(map[string]interface{})
	code := builder.code
	res := &JSend{}
	if builder.data != nil {
		data["data"] = builder.data
	}

	if code >= 200 && code < 300 {
		data["status"] = "success"
	}
	if code >= 400 && code < 500 {
		data["status"] = "fail"
	}
	if code >= 500 {
		data["status"] = "error"
	}
	if builder.message != "" {
		data["message"] = builder.message
	}
	sentBytes, err := json.Marshal(data)
	if err != nil {
		code = 500
	} else {
		res.data = sentBytes
	}
	res.code = code

	return res
}

//NewJSendBuilder ...
func NewJSendBuilder() IJSendBuilder {
	return &JSendBuilder{}
}
