package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dnguyenngoc/robot/service/settings"
)


func (ctl *Controller) TestApi(c *gin.Context){
	c.JSON(http.StatusOK, settings.Env.Server.Host)
}