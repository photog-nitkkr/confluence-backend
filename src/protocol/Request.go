package protocol

import "net/http"

type Request struct {
	Method string `json:"method,omitempty"`
	URI string `json:"uri,omitempty"`
	Params []Params `json:"params,omitempty"`
}

func GetRequestObject(r *http.Request) Request {
	method := r.Method
	uri := r.RequestURI
	var params []Params
	if method == "GET" {
		params = GetQueryParamsForGETRequest(r)
	} else {
		params = nil
	}

	return convertToRequest(method, uri, params)
}

func convertToRequest(method string, uri string, params []Params) Request {
	return Request{
		Method: method,
		URI:    uri,
		Params: params,
	}
}

