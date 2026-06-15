package utils

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func NewTokenResponse(data any, accessToken string, refreshToken string, expiresIn int) map[string]any {
	return map[string]any{
		"data": data,
		"token": Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    expiresIn,
		},
	}
}
