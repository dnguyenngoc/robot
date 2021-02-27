
package controllers

import (
	"time"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dnguyenngoc/robot/service/models"
	"github.com/dnguyenngoc/robot/service/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	helper "github.com/dnguyenngoc/robot/service/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-playground/validator/v10"
)

const (
	connectTimeout = 100
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()


// Sign up User
// @Summary Login get token by user/pass
// @Tags accounts
// @Param  account body models.UserSignUp true
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Token
// @Failure 400 {object} exceptions.BadRequest
// @Failure 500 {object} exceptions.InternalServerError
// @Router /api/v1/accounts/login/access-token [post]
func (ctl *Controller) SignUp(c *gin.Context) {

	// verify json param
	var user models.UserCreate
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	// validationErr
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	// limit time request to api attack db
	var ctx, cancel = context.WithTimeout(context.Background(), connectTimeout*time.Second)

	// verify existed or not
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error occured while checking for the email"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "this email or phone number already exists"})
		return
	}
	
	// make userDb
	var userDb models.UserDB
	userDb.First_name 			= user.First_name
	userDb.Last_name 			= user.Last_name
	userDb.Email 				= user.Email
	userDb.Password 			= helper.HashPassword(user.Password)
	userDb.Phone 				= user.Phone
	userDb.Created_at 			= helper.NowUtcTime()
	userDb.Updated_at 			= helper.NowUtcTime()
	userDb.ID 					= primitive.NewObjectID()
	userDb.User_id 				= userDb.ID.Hex()

	// handle insert to db
	_, insertErr := userCollection.InsertOne(ctx, userDb)
	defer cancel()
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": insertErr.Error()})
		return
	}
	// return Token
	token, err := helper.JwtGenerateToken(userDb.Email, userDb.First_name, userDb.Last_name, userDb.User_id)
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
// @Param  account body models.UserLogin true "Login by User/Pass" 
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Token
// @Failure 400 {object} exceptions.BadRequest
// @Failure 404 {object} exceptions.NotFound
// @Failure 500 {object} exceptions.InternalServerError
// @Router /api/v1/accounts/login/access-token [post]
func (ctl *Controller) LoginAccessToken(c *gin.Context){
	
	// verify json param
	var user models.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// validationErr
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	// limit time request to api attack db
	var ctx, cancel = context.WithTimeout(context.Background(), connectTimeout*time.Second)

	// verify existed or not
	var foundUser models.UserDB
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login or passowrd is incorrect"})
		return
	}

	// verify pass
	if helper.VerifyPassword(user.Password, foundUser.Password) != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password not match"})
		return
	}

	// return Token
	token, err := helper.JwtGenerateToken(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
	return
}

func (ctl *Controller) GetProfile(c *gin.Context){
	// verify json param
	uid := c.GetString("uid")
	c.JSON(http.StatusOK, uid)
	return
}


func (ctl *Controller) UpdateProfile(c *gin.Context){
	// verify json param
	var user *models.SignedDetails
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}




func (ctl *Controller) LoginFreshToken(c *gin.Context){
	c.JSON(200, gin.H{
		"not": "now",
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

func (ctl *Controller) DeleteAccount(c *gin.Context){
	c.JSON(200, gin.H{
		"not": "now",
	})
}