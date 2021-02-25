package controllers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/dnguyenngoc/robot/service/models"

	_ "github.com/dnguyenngoc/robot/service/exceptions"

)

// Login AccessToken
// @Description login get access-token
// @Tags accounts
// @Param  account body models.LoginAccessTokenParam true "Login by User/Pass" 
// @Accept  json
// @Produce  json
// @Success 200 {object} models.LoginAccessToken
// @Failure 400 {object} exceptions.BadRequest
// @Failure 404 {object} exceptions.NotFound
// @Failure 500 {object} exceptions.InternalServerError
// @Router /api/v1/accounts/login/access-token [post]
func (ctl *Controller) LoginAccessToken(c *gin.Context){
	c.JSON(200, gin.H{
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
		"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
    })
}