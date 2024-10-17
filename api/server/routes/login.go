package routes

import (
	"net/http"

	"github.com/JadlionHD/crud-gin-go/api/server/types"
	"github.com/JadlionHD/crud-gin-go/api/utils"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginVals types.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if loginVals.Username != "admin" && loginVals.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		c.Abort()
		return
	}

	token, err := utils.GenerateToken(123)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})

}
