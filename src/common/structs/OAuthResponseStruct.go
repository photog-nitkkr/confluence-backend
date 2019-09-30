package structs

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"net/http"
)

type OAuthResponse struct{
	IdToken string `json:"idToken,omitempty"`
}

type TokenInfo struct {
	Email string `json:"email,omitempty"`
	Sub string `json:"sub,omitempty"`
	Name string `json:"name,omitempty"`
	ProfilePic string `json:"picture,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
	College string `json:"college,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
	Year string `json:"year,omitempty"`
	OnBoard bool `json:"onBoard,omitempty"`
}

func GetIdToken(r *http.Request) (string, error) {
	var tokenStruct OAuthResponse
	err := json.NewDecoder(r.Body).Decode(&tokenStruct)

	if err != nil {
		return "", err
	}
	return tokenStruct.IdToken, nil
}

func GetTokenInfo(body []byte) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	err := json.Unmarshal(body, &tokenInfo)

	if err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func ConvertToUserObject(firestoreDocument *firestore.DocumentSnapshot, user *TokenInfo) error {
	err := firestoreDocument.DataTo(user)

	if err != nil {
		return err
	}
	return nil
}
