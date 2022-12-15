package controllers

import (
	"net/http"
	"restful-api-gin/responses"

	"github.com/gin-gonic/gin"
)

func DefaultRoute(c *gin.Context) {
	c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK,
		 Message: "Success", 
		 Data: map[string]interface{}{"data": "welcome to gin"}})

}
