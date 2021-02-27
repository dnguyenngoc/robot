/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package utils

import (
	"log"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dnguyenngoc/robot/service/settings"
	"github.com/dnguyenngoc/robot/service/models"

)

const (
	expiredAccessToken = 24
	expiredFreshToken = 48
)


func JwtGenerateToken(email string, firstName string, lastName string, uid string) (token models.Token, err error)  {
	/*
		function make jwt token. need to encrypt jwt in future
	*/
	accessClaims := &models.SignedDetails {
		Email:      email,
		FirstName: firstName,
		LastName:  lastName,
		UId:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(expiredAccessToken)).Unix(),
		},
	}

	refreshClaims := &models.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(expiredFreshToken)).Unix(),
		},
	}

	accessToken, aErr := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(settings.Env.Project.SecretKey))
	freshToken, fErr := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(settings.Env.Project.SecretKey))
	if (aErr != nil || fErr != nil) {
		log.Panic(aErr)
		return token, nil
	}
	token.AccessToken = accessToken
	token.FreshToken = freshToken
	return token, nil
}