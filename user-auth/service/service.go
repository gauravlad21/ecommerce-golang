package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gauravlad21/ecommerce-golang/user-auth/common"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceStruct) Hello(ctx context.Context) string {
	return "hello from user-auth service"
}

func (s *ServiceStruct) Signup(ctx context.Context, body *common.UserAuthBody) *common.Response {

	// validation
	if body == nil || body.Email == "" || body.Password == "" || common.StringToUserType(body.UserType) == "" {
		return common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "body is nil")
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: []string{err.Error()}}
	}

	body.Password = string(hash) // update password to hash
	_, err = s.DbOps.InsertUser(ctx, body)
	if err != nil {
		return &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: []string{err.Error()}}
	}
	return common.GetDefaultResponse()
}

func (s *ServiceStruct) Login(ctx context.Context, body *common.UserAuthBody) *common.LoginResposne {

	// validation
	if body == nil || body.Email == "" || body.Password == "" {
		return &common.LoginResposne{Response: &common.Response{StatusCode: common.StatusCode_BAD_REQUEST, ErrorMsg: []string{"body is nil"}}}
	}

	// Look up for requested user
	user, err := s.DbOps.GetUser(ctx, body.Email)
	if err != nil {
		return &common.LoginResposne{Response: &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: []string{err.Error()}}}
	}

	// Compare sent in password with saved users password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return &common.LoginResposne{Response: &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: []string{err.Error()}}}
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET")))
	if err != nil {
		return &common.LoginResposne{Response: &common.Response{StatusCode: common.StatusCode_INTERNAL_ERROR, ErrorMsg: []string{err.Error()}}}
	}
	return &common.LoginResposne{Response: &common.Response{StatusCode: common.StatusCode_OK}, Token: tokenString}
}

func (s *ServiceStruct) RequireAuth(ctx context.Context, tokenString string) (*common.User, error) {
	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString("SECRET")), nil
	})

	var ok bool
	var claims jwt.MapClaims
	if token != nil {
		if claims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			return nil, fmt.Errorf("Unautorized")
		}
	}

	// Check the expiry date
	if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
		return nil, fmt.Errorf("Unautorized")
	}

	// Find the user with token Subject
	var sub int32 = 0
	if subF64, ok := claims["sub"].(float64); ok {
		sub = int32(subF64)
	}
	user, err := s.DbOps.GetUserById(ctx, sub)
	if err != nil || user == nil || user.Email == "" {
		return nil, fmt.Errorf("Unautorized")
	}
	return user, nil
}
