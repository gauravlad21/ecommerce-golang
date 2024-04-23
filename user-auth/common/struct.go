package common

import "github.com/gauravlad21/ecommerce-golang/common"

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
	*common.Response
	Token string
}
