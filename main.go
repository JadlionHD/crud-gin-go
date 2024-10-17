package main

import (
	"net/http"

	"github.com/JadlionHD/crud-gin-go/api/server/middleware"
	"github.com/JadlionHD/crud-gin-go/api/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.POST("/login", routes.LoginHandler)
	// r.NoRoute(handle.MiddlewareFunc(), func(c *gin.Context) {
	// 	claims := jwt.ExtractClaims(c)
	// 	log.Printf("NoRoute claims: %#v\n", claims)
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	auth := router.Group("/auth", middleware.AuthMiddleware())

	auth.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "HELLO WORLD",
		})
	})

	// auth := r.Group("/auth", handle.MiddlewareFunc())
	// auth.GET("/refresh_token", handle.RefreshHandler)
	// auth.GET("/hello", func(c *gin.Context) {
	// 	claims := jwt.ExtractClaims(c)
	// 	user, _ := c.Get(middleware.IdentityKey)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"userID":   claims[middleware.IdentityKey],
	// 		"userName": user.(*types.User).UserName,
	// 		"text":     "Hello World.",
	// 	})
	// })

	router.GET("/posts/:id", routes.GetPost)
	router.POST("/posts", routes.CreatePost)
	router.DELETE("/posts/:id", routes.DeletePost)
	router.PUT("/posts/:id", routes.UpdatePost)

	router.Run("localhost:8080")
}
