package util

import "time"

func GetCurrentDate() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
