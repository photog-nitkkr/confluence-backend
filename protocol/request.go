package protocol

import (
	"encoding/json"
	"net/http"
)

type (
	RequestInterface interface {
		GetMethod() string
		GetURI() string
	}

	Request struct {
		Method string
		URI string
	}
)
//getters
func (R *Request) GetMethod() string {
	return R.Method
}

func (R *Request) GetURI() string {
	return R.URI
}

//setters
func SetMethod(r *Request, method string) *Request {
	r.Method = method
	return r
}

func SetURI(r *Request, uri string) *Request {
	r.URI = uri
	return r
}

//extra functionality
func GiveRequestObject(r *http.Request) *Request {
	return &Request{
		Method: r.Method,
		URI:    r.RequestURI,
	}
}

func GiveRequestInJson(r *http.Request) string {
	request := GiveRequestObject(r)

	jsonObject, err := json.Marshal(request)

	if err == nil {
		return string(jsonObject)
	} else {
		return "{}"
	}
}
