package model

type TodoRequest struct {
	Content      string `json:"content" binding:"required"`
	Completed    bool   `json:"completed"`
	UserId       int64  `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
