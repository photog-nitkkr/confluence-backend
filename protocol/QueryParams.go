package protocol

import "net/http"

type Params struct {
	Key string
	Value []string
}

func GetQueryParamsForGETRequest(r *http.Request) []Params {
	queryMap := r.URL.Query()
	var params []Params
	for key, val := range queryMap {
		param := getQueryParam(val, key)
		params = append(params, param)
	}
	return params
}

func getQueryParam(val []string, key string) Params {
	var valueArray []string

	for _, indexValue := range val {
		valueArray = append(valueArray, indexValue)
	}

	return convertToParams(key, valueArray)
}

func convertToParams(key string, val []string) Params {
	return Params{
		Key:   key,
		Value: val,
	}
}
