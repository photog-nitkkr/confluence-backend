package sponsorRoutes

import (
	. "net/http"
	"protocol"

	. "sponsor"
)

func readSponsor(w ResponseWriter, r *Request) {
	if len(r.URL.Query()) == 0 {
		returnSponsors(w, r)
		return
	} else {
		returnInvalidParamsError(w, r)
		return
	}
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

func returnSponsors(w ResponseWriter, r *Request) {
	sponsors, err := GetSponsorsFromFirestore()

	if err != nil {
		responseObject := protocol.Response{
			Success: false,
			Message: "Error in Getting Sponsors",
			Request: protocol.GetRequestObject(r),
			Data:    nil,
		}
		protocol.WriteResponseObject(w, r, responseObject, StatusInternalServerError)
		return
	}

	responseObject := protocol.Response{
		Success: true,
		Message: "Giving All Sponsors",
		Request: protocol.GetRequestObject(r),
		Data:    *sponsors,
	}
	protocol.WriteResponseObject(w, r, responseObject, StatusOK)
	return
}
