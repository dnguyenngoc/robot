/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const salt = 14

func VerifyPassword(userPassword string, providedPassword string) (bool) {
    err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
    check := true
    if err != nil {
        check = false
    }

    return check
}


func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
    if err != nil {
        log.Panic(err)
    }

    return string(bytes)
}
