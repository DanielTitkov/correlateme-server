package model

type (
	CreateUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	GetTokenRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	GetTokenResponse struct {
		Token string `json:"token"`
	}

	GetUserRequest struct{}

	GetUserResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
