package main

import (
	"net/http"

	"github.com/LilzBay/go-jwt/controllers"
	"github.com/LilzBay/go-jwt/initializers"
	"github.com/LilzBay/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVarialble()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":  "lin",
			"lover": "zao",
		})
	})

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
