package eventNames

import (
	"../../common"
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

		common.WriteResponseOld(w, jsonObject, http.StatusInternalServerError)
		return
	} else {
		jsonObj, errInJson := json.Marshal(eventNames)

		if errInJson != nil {
			jsonObject := protocol.GiveResponseJsonUsingArguments(false,
				"Error in converting to Json",
				protocol.GiveRequestInJson(r),
				"",
			)

			common.WriteResponseOld(w, jsonObject, http.StatusInternalServerError)
			return
		}

		jsonObject := protocol.GiveResponseJsonUsingArguments(true,
			"Events Names by All categoryRoutes: ",
			protocol.GiveRequestInJson(r),
			string(jsonObj),
		)

		common.WriteResponseOld(w, jsonObject, http.StatusOK)
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

		common.WriteResponseOld(w, jsonObject, http.StatusInternalServerError)
		return
	} else {
		jsonObj, errInJson := json.Marshal(array)

		if errInJson != nil {
			jsonObject := protocol.GiveResponseJsonUsingArguments(false,
				"Error in converting to Json",
				protocol.GiveRequestInJson(r),
				"",
			)

			common.WriteResponseOld(w, jsonObject, http.StatusInternalServerError)
			return
		}

		jsonObject := protocol.GiveResponseJsonUsingArguments(true,
				"Events Names by category: "+category,
				protocol.GiveRequestInJson(r),
				string(jsonObj),
		)

		common.WriteResponseOld(w, jsonObject, http.StatusOK)
		return
	}
	return
}

