package protocol

import "net/http"

type Request struct {
	Method string
	URI string
	Params []Params
}

func GetRequestObject(r *http.Request) Request {
	method := r.Method
	uri := r.RequestURI
	params := GetQueryParams(r)

	return convertToRequest(method, uri, params)
}

func convertToRequest(method string, uri string, params []Params) Request {
	return Request{
		Method: method,
		URI:    uri,
		Params: params,
	}
}

