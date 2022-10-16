package util

import (
	"strconv"
	"time"
)

func GetCurrentTimeStamp() (timestamp string) {
    return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}
