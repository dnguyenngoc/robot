package controllers

import (
	// "context"
	// "time"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/dnguyenngoc/robot/service/models"
	"github.com/dnguyenngoc/robot/service/db_models"
	_ "github.com/dnguyenngoc/robot/service/exceptions"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// helper "github.com/dnguyenngoc/robot/service/utils"
)


// Sign up User
// @Summary Login get token by user/pass
// @Tags accounts
// @Param  account body models.UserSignUp true
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserLogin
// @Failure 400 {object} exceptions.BadRequest
// @Failure 404 {object} exceptions.NotFound
// @Failure 500 {object} exceptions.InternalServerError
// @Router /api/v1/accounts/login/access-token [post]
func (ctl *Controller) SignUp() gin.HandlerFunc {
	return func(c *gin.Context){
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user db_models.User
		// if err := c.BindJSON(&user); err != nil {
        //     c.JSON(http.StatusBadRequest, exceptions.BadRequest})
        //     return
        // }
		// user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		// user.ID = primitive.NewObjectID()
		// user.User_id = user.ID.Hex()
		// acessToken, refreshToken, _ := helper.JwtGenerateToken(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		// user.Access_token = acessToken
		// user.Refresh_token = refreshToken
		// resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		// if insertErr != nil {
		// 	msg := fmt.Sprintf("User item was not created")
		// 	c.JSON(http.StatusInternalServerError, exceptions.InternalServerError)
		// 	return
		// }
		c.JSON(http.StatusOK, user)
	}
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