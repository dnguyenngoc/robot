package main

import (
	"github.com/gin-gonic/gin"

	"app/config"

	"app/routes"

	"log"

	"strconv"
)

func init() {

	// Load env like config.Variable for all system can call it (-.0)!!
	config.ViperReadEnvPath(".", "config.yaml", "yaml")

}

func main() {

	env := config.GetEnv()

	gin.SetMode(gin.ReleaseMode)

	router := routes.SetupRoutes()

	port := env.Server.Port

	log.Println("Server starting on port: ", port)

	if err := router.Run(":" + strconv(port)); err != nil {
		log.Panicf("error: %s", err)
	}

}
