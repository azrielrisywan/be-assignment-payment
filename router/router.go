// router/router.go

package router

import (
	"azrielrisywan/be-assignment-payment/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Testing routes
	router.GET("/test", controller.Test)

	router.POST("/testPost", controller.TestPostRequest)

	// Supabase routes
	router.POST("/signup", controller.SignUp)

	router.POST("/signin", controller.SignIn)

	// Account routes
	router.POST("/getAccountsByUser", controller.GetAccountsByUser)

	return router
}
