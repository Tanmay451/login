package routes

import (
	"login/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

//SetupRouter sets up routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(os.Getenv("GIN_MODE"))

	router.GET("/", controllers.LogIN)

	// router.NoRoute(controllers.PageNotFound)

	return router
}
