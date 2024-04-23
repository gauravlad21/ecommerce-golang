package common

type UserAuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type User struct {
	ID       int32
	Email    string
	Password string
}

type LoginResposne struct {
	*Response
	Token string
}
