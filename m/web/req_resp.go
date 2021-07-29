package web

// ******************************
// error response

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(s ...string) ErrorResponse {
	if len(s) == 0 {
		return ErrorResponse{Error: "param error"}
	}
	return ErrorResponse{Error: s[0]}
}

// ******************************
// oauth

type OAuthRequest struct {
}

type OAuthResponse struct {
	UserId string `json:"userId"`
}
