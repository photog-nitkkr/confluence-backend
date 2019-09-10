package eventNameRoutes

import (
	"../../event"
	"../../protocol"
	. "net/http"
)

func readCategory(w ResponseWriter, r *Request) {
	category := r.URL.Query()["category"]
	event := r.URL.Query()["event"]
	if len(r.URL.Query()) == 0 {
		return
	} else if category != nil && event == nil {
		return
	} else if category != nil && event != nil && len(category) == 1 {
		returnEvent(w, r, category[0], event)
	} else {
		returnInvalidParamsError(w, r)
		return
	}
}

func returnEvent(w ResponseWriter, r *Request, categoryName string, eventsName []string) {
	events, err := event.GetEventsName(categoryName, eventsName)

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: err.Error(),
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving data for given events",
		Request: protocol.GetRequestObject(r),
		Data:    events,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func returnInvalidParamsError(w ResponseWriter, r *Request) {
	responseObject := protocol.Response{
		Success: false,
		Message: "Invalid Parameters",
		Request: protocol.GetRequestObject(r),
		Data:    nil,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusBadGateway)
	return
}

