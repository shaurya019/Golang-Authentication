package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shaurya019/controllers"
	middleware "github.com/shaurya019/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}