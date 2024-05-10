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

	return router
}
