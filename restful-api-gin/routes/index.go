package routes

import (
	"restful-api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine){
	router.GET("/",controllers.DefaultRoute)
}
