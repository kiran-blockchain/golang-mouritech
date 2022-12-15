package main

import (
	"fmt"
	config "restful-api-gin/config"
	"restful-api-gin/database"
	routes "restful-api-gin/routes"

	"github.com/gin-gonic/gin"
)
func main(){
//step 1 - get the router instance from Gin
//step 2 - create a default route and route handler
//step 3 - run the app on the port

    router:= gin.Default()
	routes.AppRoutes(router)
    database.DB = database.ConnectDB()
	fmt.Println("Server running on port:",config.GetEnvData("PORT"))
	router.Run(config.GetEnvData("PORT"))
}
	
//Step -2 
