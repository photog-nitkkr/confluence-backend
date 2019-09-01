package eventNames

import (
	"../../protocol"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//EventNamesRouteHandler
func EventNamesRouteHandler() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", getEventNames).Methods("GET")
	return muxRouter
}



func getEventNames(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	if r.URL.Query()["category"] != nil {
		getEventNameByCategory(w, r)
		return
	} else {
		getAllEventNames(w, r)
		return
	}

}

func getAllEventNames(w http.ResponseWriter, r *http.Request) {
	eventNames, err := GetAllEventName()

	if err != nil {
		jsonObject := protocol.GiveResponseJsonUsingArguments(
			false,
			"Error in getting events",
			protocol.GiveRequestInJson(r),
			err.Error(),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(jsonObject))
		return
	} else {
		jsonObj, errInJson := json.Marshal(eventNames)

		if errInJson != nil {
			jsonObject := protocol.GiveResponseJsonUsingArguments(false,
				"Error in converting to Json",
				protocol.GiveRequestInJson(r),
				"",
			)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(jsonObject))
			return
		}

		jsonObject := protocol.GiveResponseJsonUsingArguments(true,
			"Events Names by All category: ",
			protocol.GiveRequestInJson(r),
			string(jsonObj),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonObject))
		return
	}
	return

}

func getEventNameByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query()["category"][0]
	array, err := GetEventByCategory(category)

	if err != nil {
		jsonObject := protocol.GiveResponseJsonUsingArguments(
			false,
			"Error in getting events",
			protocol.GiveRequestInJson(r),
			err.Error(),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(jsonObject))
		return
	} else {
		jsonObj, errInJson := json.Marshal(array)

		if errInJson != nil {
			jsonObject := protocol.GiveResponseJsonUsingArguments(false,
				"Error in converting to Json",
				protocol.GiveRequestInJson(r),
				"",
			)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(jsonObject))
			return
		}

		jsonObject := protocol.GiveResponseJsonUsingArguments(true,
				"Events Names by category: "+category,
				protocol.GiveRequestInJson(r),
				string(jsonObj),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonObject))
		return
	}
	return
}

