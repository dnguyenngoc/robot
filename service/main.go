package main

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/dnguyenngoc/robot/service/routes"
	"github.com/dnguyenngoc/robot/service/settings"
	"github.com/dnguyenngoc/robot/service/database"
	"github.com/dnguyenngoc/robot/service/docs" 
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/files" // swagger embed files
)

func init() {

	log.Println("Starting server...")

	// Set gin mode
	gin.SetMode(gin.ReleaseMode)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "API DOCUMENTATION"
	docs.SwaggerInfo.Description = "This is a go server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "robot.com"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	log.Println("API Documentation at /swagger/index.html")

	// make config file
	settings.ViperReadEnvPath("./settings", "variable.yaml", "yaml")

	// initial db
	database.DBinstance()

	// log.Println("Load all env to conf structure")

}

func main() {

	// load config
	conf := settings.GetEnv()

	// load router config
	router := routes.SetupRoutes()

	// load swaggger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	log.Println("Server starting on port: ", conf.Server.Port)

	if err := router.Run(":" + strconv.Itoa(conf.Server.Port)); err != nil {
		log.Panicf("error: %s", err)
	}

}
