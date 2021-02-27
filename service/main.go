/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package main

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/dnguyenngoc/robot/service/routes"
	"github.com/dnguyenngoc/robot/service/settings"
	"github.com/dnguyenngoc/robot/service/docs" 
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	"github.com/dnguyenngoc/robot/service/database"
)

func init(){
	/*
        Loading all component for service include gin, mongodb, router api, swagger, ...
    */

	log.Println("Starting server...")

	// Set gin mode
	gin.SetMode(gin.ReleaseMode)
	log.Println("Seting Gin Mode: Done!")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "API DOCUMENTATION"
	docs.SwaggerInfo.Description = "This is a go server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "robot.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	log.Println("API Documentation at /swagger/index.html")

	// make config file
	settings.InitialViperWriteCofig("./settings", "development.yaml", "yaml")
	log.Println("Load environment variable: Good!")		

	// connect db ping
	database.DBinstance()
	log.Println("Connected and pinged mongodb: Successfully!")
}

func main() {
	/*
        Loading all component for service include gin, mongodb, router api, swagger, ...
    */
	
	// load router config
	router := routes.SetupRoutes()
	log.Println("Setup Routes: Completed!")

	// load swaggger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Server starting on port: ", settings.Env.Server.Port)

	// run service
	if err := router.Run(":" + strconv.Itoa(settings.Env.Server.Port)); err != nil {
		log.Panicf("error: %s", err)
	}



}
