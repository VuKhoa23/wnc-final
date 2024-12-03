package model

// mapping for request, response

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
