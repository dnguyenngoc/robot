package controllers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/dnguyenngoc/robot/service/models"

	_ "github.com/dnguyenngoc/robot/service/exceptions"

)

// Login AccessToken
// @Summary Login get token by user/pass
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

func (ctl *Controller) LoginFreshToken(c *gin.Context){
	c.JSON(200, gin.H{
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
		"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
    })
}

func (ctl *Controller) Logout(c *gin.Context){
	c.JSON(200, gin.H{
		"not": "now",
	})
}

func (ctl *Controller) CreateAccount(c *gin.Context){
	c.JSON(200, gin.H{
		"not": "now",
	})
}

func (ctl *Controller) UpdateProfile(c *gin.Context){
	c.JSON(200, gin.H{
		"not": "now",
	})
}

func (ctl *Controller) DeleteAccount(c *gin.Context){
	c.JSON(200, gin.H{
		"not": "now",
	})
}