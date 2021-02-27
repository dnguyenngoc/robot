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

var conf settings.Config = settings.GetEnv()

const (
	expiredAccessToken = 24
	expiredFreshToken = 48
)

func JwtGenerateToken(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error)  {
	accessClaims := &models.SignedDetails{
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

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(conf.Project.SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(conf.Project.SecretKey))

	if err != nil {
		log.Panic(err)
		return
	}

	return accessToken, refreshToken, err
}