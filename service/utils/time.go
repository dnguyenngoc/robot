/* 
	@Author: Duy Nguyen
	@Email: <duynguyenngoc@hotmail.com>
*/

package utils

import "time"


func  NowUtcTime() time.Time {
	/*
		Get now utc time
	*/
	loc, _ := time.LoadLocation("UTC")
    now := time.Now().In(loc)
	return now
}
