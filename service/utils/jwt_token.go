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
	"fmt"
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
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
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


func ValidateToken(signedToken string) (claims *models.SignedDetails , msg string) {
	/*
		ValidateToken validates the jwt token
	*/
    token, err := jwt.ParseWithClaims(
        signedToken,
        &models.SignedDetails{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte(settings.Env.Project.SecretKey), nil
        },
    )

    if err != nil {
        msg = err.Error()
        return
    }

    claims, ok := token.Claims.(*models.SignedDetails)
    if !ok {
        msg = fmt.Sprintf("the token is invalid")
        msg = err.Error()
        return
    }

    if claims.ExpiresAt < time.Now().Local().Unix() {
        msg = fmt.Sprintf("token is expired")
        msg = err.Error()
        return
    }

    return claims, msg
}