package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	body := &common.AuthorizationTokenRequest{Token: tokenString}
	data, _ := json.Marshal(body)
	byteRes, err := common.HttpCall(context.Background(), "POST", "http://localhost:5001/user-auth//authorize", nil, nil, data)
	if err != nil {
		return
	}

	res := &common.AuthorizationTokenResponse{}
	json.Unmarshal(byteRes, res)
	if res != nil && res.IsAuthorized {
		c.Set("user", res.Email)
		c.Next()
	}
}
