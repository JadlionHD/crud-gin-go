package middleware

import (
	"fmt"
	"net/http"

	"github.com/JadlionHD/crud-gin-go/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("token")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		jwtToken, err := utils.TokenValidate(token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)

		if !ok && !jwtToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		fmt.Println("userid: ", claims["user_id"])
		c.Next()
	}
}
