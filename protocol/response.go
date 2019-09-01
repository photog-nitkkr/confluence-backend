package protocol

import "encoding/json"

type (
	ResponseInterface interface {
		GetSuccess() bool
		GetMessage() string
		GetRequest() string
		GetData() string
	}

	Response struct {
		Success bool
		Message string
		Request string
		Data string
	}
)

//getters
func (r *Response) GetSuccess() bool {
	return r.Success
}

func (r *Response) GetMessage() string {
	return r.Message
}

func (r *Response) GetRequest() string {
	return r.Request
}

func (r *Response) GetData() string {
	return r.Data
}

//setters
func SetSuccess(r *Response, success bool) *Response {
	r.Success = success
	return r
}

func SetMessage(r *Response, message string) *Response {
	r.Message = message
	return r
}

func SetRequest(r *Response, request string) *Response {
	r.Request = request
	return r
}

func SetData(r *Response, data string) *Response {
	r.Data = data
	return r
}


//extra functionality
func GiveResponseObject(success bool, message string, request string, data string) *Response {
	return &Response{
		Success: success,
		Message: message,
		Request: request,
		Data:    data,
	}
}

func GiveResponseJsonUsingResponse(r *Response) string {
	object, err := json.Marshal(r)
	if err != nil {
		return "{}"
	} else {
		return string(object)
	}
}

func GiveResponseJsonUsingArguments(success bool, message string, request string, data string) string {
	return GiveResponseJsonUsingResponse(GiveResponseObject(success, message, request, data))
}




