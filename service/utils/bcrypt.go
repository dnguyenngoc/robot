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
    /*
        Verify hash pass on db and clearly pass
    */
    err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
    check := true
    if err != nil {
        check = false
    }

    return check
}


func HashPassword(password string) string {
    /*
        encrypt password
    */
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
    if err != nil {
        log.Panic(err)
    }

    return string(bytes)
}
