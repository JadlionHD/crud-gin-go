package middleware

import (
	"net/http"
	"time"

	"github.com/JadlionHD/crud-gin-go/api/server/types"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	IdentityKey = "id"
)

func JWTInitParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKey,
		PayloadFunc:     payloadFunc(),
		IdentityHandler: identityHandler(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		Unauthorized:    unauthorized(),
		// TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		SendCookie:     true,
		SecureCookie:   false,
		CookieHTTPOnly: true,
		// CookieDomain:   "localhost:8080",
		CookieName:     "token",
		CookieSameSite: http.SameSiteDefaultMode,
		// TokenLookup: "query:token",
		TokenLookup: "cookie:token",
		// TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*types.User); ok {
			return jwt.MapClaims{
				IdentityKey: v.UserName,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &types.User{
			UserName: claims[IdentityKey].(string),
		}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals types.Login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
			return &types.User{
				UserName:  userID,
				FirstName: "Hendra",
				LastName:  "Gunawan",
			}, nil
		}
		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*types.User); ok && v.UserName == "admin" {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"message": message,
		})
	}
}
