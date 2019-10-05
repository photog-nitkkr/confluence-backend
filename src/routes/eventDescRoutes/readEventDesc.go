package eventDescRoutes

import (
	"event"
	. "net/http"
	"protocol"
)

func readEventDesc(w ResponseWriter, r *Request) {
	category := r.URL.Query()["category"]
	event := r.URL.Query()["event"]

	if len(r.URL.Query()) == 0 {
		returnAllEvents(w, r)
		return
	} else if category != nil && len(event) == 0 {
		returnCategoryEvents(w, r, category)
		return
	} else if category != nil && len(category) == 1 {
		returnEvent(w, r, category[0], event)
		return
	} else if category == nil && event != nil && len(event) == 1 {
		returnSpecificEvent(w, r, event[0])
	} else {
		returnInvalidParamsError(w, r)
		return
	}
}

func returnSpecificEvent(w ResponseWriter, r *Request, eventName string) {
	event, err := event.GetSpecificEvent(eventName, "eventDesc")

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
		Message: "Giving required Event",
		Request: protocol.GetRequestObject(r),
		Data:    event,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func returnAllEvents(w ResponseWriter, r *Request) {
	category, err := event.GetAllEventsForAllCategory("eventDesc")

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
		Message: "Giving events for all categories",
		Request: protocol.GetRequestObject(r),
		Data:    category,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func returnCategoryEvents(w ResponseWriter, r *Request, categoryName []string) {
	categories, err := event.GetEventsForCategories(categoryName, "eventDesc")

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
		Message: "Giving events for given categories",
		Request: protocol.GetRequestObject(r),
		Data:    categories,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}

func returnEvent(w ResponseWriter, r *Request, categoryName string, eventsName []string) {
	events, err := event.GetEvents(categoryName, eventsName, "eventDesc")

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

