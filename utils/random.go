package utils

import (
	"strconv"
	"time"
)

//GenerateRandomFileName -> generates the fileName with unique time
func GenerateRandomFileName() string {
	time := time.Now().UnixNano()
	return strconv.FormatInt(time, 10)
}
