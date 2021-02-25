
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dnguyenngoc/robot/service/controllers"

)
func SetupRoutes() *gin.Engine {

	// Define router
	router := gin.Default()

	// Define controller
	control := controllers.NewController()

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

	// Routes to v1
	api := router.Group("api")
	v1 := api.Group("v1")

	// Routes account
	account := v1.Group("accounts")
	account.POST("/login/access-token", control.LoginAccessToken)
	account.POST("/login/refresh-token", control.LoginAccessToken)
	account.POST("/logout", control.LoginAccessToken)
	account.POST("/new", control.LoginAccessToken)
	account.PUT("/profile", control.LoginAccessToken)
	account.DELETE("", control.LoginAccessToken)

	return router
}
