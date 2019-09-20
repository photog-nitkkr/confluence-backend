package structs

import (
	"encoding/json"
	"net/http"
)

type OAuthResponse struct{
	IdToken string `json:"id_token"`
}

type TokenInfo struct {
	Email string `json:"email"`
	Sub string `json:"sub"`
	Name string `json:"name"`
	ProfilePic string `json:"picture"`
	ErrorDescription string `json:"error_description"`
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