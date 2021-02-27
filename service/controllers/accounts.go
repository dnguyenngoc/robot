
package controllers

import (
	// "context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dnguyenngoc/robot/service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	helper "github.com/dnguyenngoc/robot/service/utils"
)

// Sign up User
// @Summary Login get token by user/pass
// @Tags accounts
// @Param  account body models.UserSignUp true
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserCreate
// @Failure 400 {object} exceptions.BadRequest
// @Failure 404 {object} exceptions.NotFound
// @Failure 500 {object} exceptions.InternalServerError
// @Router /api/v1/accounts/login/access-token [post]
func (ctl *Controller) SignUp(c *gin.Context) {

	// limit time request to api
	// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	
	// verify json param
	var user models.UserCreate
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// verify existed or not


	// make Time
	user.CreatedAt = helper.NowUtcTime()
	user.UpdatedAt = helper.NowUtcTime()
	
	// make userDb
	var userDb models.UserDB
	userDb.UserCreate = user
	userDb.ID = primitive.NewObjectID()
	userDb.UserId = userDb.ID.Hex()

	// handle insert to db
	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
			c.JSON(http.StatusInternalServer, gin.H{"message": insertErr.Error()})
	}	

	// return Token
	token, err := helper.JwtGenerateToken(*userDb.Email, *userDb.FirstName, *userDb.LastName, userDb.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
	return
}


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