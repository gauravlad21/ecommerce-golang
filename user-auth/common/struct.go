package common

type UserAuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type User struct {
	ID       int32  `json:"id"`
	Email    string `json:"email"`
	Password string
	UserType string `json:"user_type"`
}

type LoginResposne struct {
	*Response
	Token string
}

type AuthorizationTokenRequest struct {
	Token string `json:"token"`
}

type AuthorizationTokenResponse struct {
	IsAuthorized bool  `json:"is_authorized"`
	User         *User `json:"user"`
}
