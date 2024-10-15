package main

import (
	"log"
	"net/http"

	"github.com/JadlionHD/crud-gin-go/api/server/middleware"
	"github.com/JadlionHD/crud-gin-go/api/server/routes"
	"github.com/JadlionHD/crud-gin-go/api/server/types"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authMiddleware, err := jwt.New(middleware.JWTInitParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.Use(handlerMiddleWare(authMiddleware))
	registerRoute(router, authMiddleware)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	router.GET("/posts/:id", routes.GetPost)
	router.POST("/posts", routes.CreatePost)
	router.DELETE("/posts/:id", routes.DeletePost)
	router.PUT("/posts/:id", routes.UpdatePost)

	router.Run("localhost:8080")
}

func registerRoute(r *gin.Engine, handle *jwt.GinJWTMiddleware) {
	r.POST("/login", handle.LoginHandler)
	r.NoRoute(handle.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth", handle.MiddlewareFunc())
	auth.GET("/refresh_token", handle.RefreshHandler)
	auth.GET("/hello", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		user, _ := c.Get(middleware.IdentityKey)
		c.JSON(http.StatusOK, gin.H{
			"userID":   claims[middleware.IdentityKey],
			"userName": user.(*types.User).UserName,
			"text":     "Hello World.",
		})
	})
}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}
