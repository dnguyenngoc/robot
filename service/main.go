package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dnguyenngoc/robot/service/routes"

	"github.com/dnguyenngoc/robot/service/settings"

	"log"

)

func init() {
	config.ViperReadEnvPath("./settings", "variable.yaml", "yaml")
}

func main() {

	log.Println("Starting server...")

	gin.SetMode(gin.ReleaseMode)

	router := routes.SetupRoutes()

	port := "8080"

	log.Println("Server starting on port: ", port)

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
