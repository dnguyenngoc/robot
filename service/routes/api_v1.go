
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/swaggo/files" // swagger embed files

	"github.com/dnguyenngoc/robot/service/controllers"	

)
func SetupRoutes() *gin.Engine {

	// Define router
	router := gin.Default()

	// Define controller
	ctl := controllers.NewController()

	// load template
	router.LoadHTMLGlob("template/*")

	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Routes to api
	api := router.Group("api")

	// Routes to v1
	v1 := api.Group("v1")

	// Routes account
	account := v1.Group("accounts")

	account.GET("/login/access-token", ctl.LoginAccessToken)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
