package common

import (
	"../protocol"
	"encoding/json"
	"net/http"
)

func WriteResponseOld(w http.ResponseWriter, response string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
	return
}

func writeResponse(w http.ResponseWriter, response string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
	return
}

func WriteResponseObject(w http.ResponseWriter, responseObject protocol.ResponseProtocol, statusCode int) {
	jsonResponse, err := json.Marshal(responseObject)

	if err != nil {
		return
	} else {
		writeResponse(w, string(jsonResponse), statusCode)
	}
}
