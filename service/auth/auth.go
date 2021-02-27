/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	helper "github.com/dnguyenngoc/robot/service/utils"
	"strings"

)


func Authentication() gin.HandlerFunc {
	/*
		Handle jwt token with status: Unauthorized, authorized and expire
	*/
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if strings.HasPrefix(clientToken, "Bearer ") == false {
			c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("Authorization needed start with Bearer")})
			c.Abort()
			return
		} else {
			splitToken := strings.Split(clientToken, "Bearer ")
			claims, err := helper.ValidateToken(splitToken[1])
			if err != "" {
				c.JSON(http.StatusUnauthorized, gin.H{"message": err})
				c.Abort()
				return
			}
			c.Set("email", claims.Email)
			c.Set("first_name", claims.First_name)
			c.Set("last_name", claims.Last_name)
			c.Set("uid", claims.Uid)
			c.Set("user_type", claims.User_type)

			c.Next()
		}		
	}
}