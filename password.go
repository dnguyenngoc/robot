/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package utils

import (
	"log"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) string {
	/*
		hasing string password by bcrypt
	*/

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

    if err != nil {
        log.Panic(err)
    }

    return string(bytes)
}


func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	/*
		verify hasing password by bcrypt
	*/

    err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
    check := true
    msg := ""

    if err != nil {
        msg = fmt.Sprintf("login or passowrd is incorrect")
        check = false
    }

    return check, msg
}