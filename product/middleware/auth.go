package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gin-gonic/gin"
)

func IsAutheniticated(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token := strings.Split(tokenString, " ")[1]

	body := &common.AuthorizationTokenRequest{Token: token}
	data, err := json.Marshal(body)
	if err != nil {
		// fmt.Print("lol1.1\n", err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	byteRes, err := common.HttpCall(context.Background(), "POST", "http://localhost:5001/user-auth/authorize", nil, nil, data)
	if err != nil {
		// fmt.Print("lol2\n")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res := &common.AuthorizationTokenResponse{}
	err = json.Unmarshal(byteRes, res)
	if err != nil {
		// fmt.Print("lol2.5\n", err.Error(), string(byteRes))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Print(res.IsAuthorized)
	if res != nil && res.IsAuthorized && res.User != nil {
		c.Set("user", res.User)
		c.Next()
	} else {
		// fmt.Print("lol3\n")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
