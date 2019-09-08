package protocol

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Request Request `json:"request"`
	Data interface{} `json:"data"`
}

func WriteResponseObject(w http.ResponseWriter, r *http.Request, responseObject Response, statusCode int) {
	jsonResponse, err := json.Marshal(responseObject)

	if err != nil {
		writeJsonError(w, r)
		return
	} else {
		writeResponse(w, string(jsonResponse), statusCode)
		return
	}
}

func writeJsonError(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: false,
		Message: "Error In Json Conversion",
		Request: GetRequestObject(r),
		Data:    nil,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		writeResponse(w, "", http.StatusBadGateway)
		panic(err)
		return
	} else {
		writeResponse(w, string(jsonResponse), http.StatusInternalServerError)
		return
	}
}


func writeResponse(w http.ResponseWriter, response string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
	return
}