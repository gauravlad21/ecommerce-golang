package controller

import (
	"net/http"
	"strings"

	"github.com/gauravlad21/ecommerce-golang/user-auth/common"
	"github.com/gin-gonic/gin"
)

func Hello(oldctx *gin.Context) {
	ctx := common.GetContext(oldctx)
	msg := serviceRepo.Hello(ctx)
	oldctx.JSON(200, msg)
}

func Signup(c *gin.Context) {
	body := &common.UserAuthBody{}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	res := serviceRepo.Signup(common.GetContext(c), body)
	c.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	body := common.UserAuthBody{}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}

	res := serviceRepo.Login(common.GetContext(c), &body)
	if res != nil && res.StatusCode == common.StatusCode_OK {
		// c.SetSameSite(http.SameSiteLaxMode)
		// c.SetCookie("Authorization", res.Token, 60*5, "", "", false, true) // todo
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{})
}

// kind of middleware, but it uses database
func RequireAuth(c *gin.Context) {
	// tokenString, err := c.Cookie("Authorization")
	splitToken := strings.Split(c.Request.Header["Authorization"][0], " ")
	if len(splitToken) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	tokenString := splitToken[1]

	user, err := serviceRepo.RequireAuth(common.GetContext(c), tokenString)
	if err != nil {
		// fmt.Printf("error3: %v\n", err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Attach the request
	c.Set("user", user)
	//Continue
	c.Next()
}

func Authorized(c *gin.Context) {
	body := &common.AuthorizationTokenRequest{}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, common.GetErrMsgsResponse(common.StatusCode_BAD_REQUEST, "Failed to read body"))
		return
	}
	tokenString := body.Token
	var res *common.AuthorizationTokenResponse
	user, err := serviceRepo.RequireAuth(common.GetContext(c), tokenString)
	if err != nil || user.Email == "" {
		res = &common.AuthorizationTokenResponse{IsAuthorized: false}
	} else {
		res = &common.AuthorizationTokenResponse{IsAuthorized: true, User: user}
	}
	c.JSON(http.StatusOK, res)
}

func Validate(c *gin.Context) {
	contectUser, _ := c.Get("user")

	user := contectUser.(*common.User)
	user.Password = "updated from handler"
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
