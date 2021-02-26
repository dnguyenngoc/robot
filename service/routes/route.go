/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dnguyenngoc/robot/service/controllers"
)


func SetupRoutes() *gin.Engine {
	/*
        Setup Router for this system. 
		Include router (using gin.Default() it mean using enable middleware)
    */

	// Define router
	router := gin.Default()

	// Define controller
	control := controllers.NewController()

	// load template
	router.LoadHTMLGlob("templates/*")

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
	api := router.Group("api"); 
	{
		v1 := api.Group("v1"); 
		{
			// Routes account
			account := v1.Group("accounts"); 
			{
				account.POST("/login/access-token", control.LoginAccessToken)
				account.POST("/login/refresh-token", control.LoginFreshToken)
				account.POST("/logout", control.Logout)
				account.POST("/new", control.CreateAccount)
				account.PUT("/profile", control.UpdateProfile)
				account.DELETE("", control.DeleteAccount)
			}	
		}
	}
	
	return router
}
