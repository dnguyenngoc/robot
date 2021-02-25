package main

import (

	"log"

	"github.com/gin-gonic/gin"

	"github.com/dnguyenngoc/robot/service/routes"

	"github.com/dnguyenngoc/robot/service/settings"

	"github.com/dnguyenngoc/robot/service/docs" 

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	swaggerFiles "github.com/swaggo/files" // swagger embed files

)

func init() {
	config.ViperReadEnvPath("./settings", "variable.yaml", "yaml")
}

func main() {
	
	log.Println("Starting server...")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "API DOCUMENTATION"
	docs.SwaggerInfo.Description = "This is a go server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "robot.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	log.Println("API Documentation at /swagger/index.html")

	// Set gin mode
	gin.SetMode(gin.ReleaseMode)

	// load router config
	router := routes.SetupRoutes()

	// load swaggger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	log.Println("Server starting on port: ", "8080")

	if err := router.Run(":" + "8080"); err != nil {
		log.Panicf("error: %s", err)
	}
}
