package authRoutes

import (
	"fmt"
	. "net/http"
	"protocol"
)

func login(w ResponseWriter, r *Request) {
	getUserEmail(r)
}


func getUserEmail(r *Request) {
	googleUrl := "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token="
	idToken, err := protocol.GetSpecificTagFromBody("id_token", r)
	fmt.Println(idToken)
	fmt.Println(err)
	idTokenString := fmt.Sprintf("%v", idToken)

	finalUrl := googleUrl + string(idTokenString)
	fmt.Println(idTokenString)
	resp, _ := Get(finalUrl)

	fmt.Println(resp)
}
