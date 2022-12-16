package routes

import (
	"restful-api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine){
	router.POST("/create",controllers.CreateUser())
}
