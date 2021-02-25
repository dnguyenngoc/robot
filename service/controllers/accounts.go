package controllers

import "github.com/gin-gonic/gin"

// ShowAccount godoc
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/accounts/login/access-token [get]
func (ctl *Controller) LoginAccessToken(c *gin.Context){
	c.JSON(200, gin.H{
        "access_token": "fake_access_token",
		"expire_token": "50000",
		"refresh_token": "fake_refresh_token",
    })
}