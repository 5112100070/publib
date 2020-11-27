package time

import "time"

func GetTimeWIB() time.Time {
	location, _ := time.LoadLocation("Asia/Jakarta")
	wib := time.Now().In(location)
	return wib
}
