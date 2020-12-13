package date_utils

import (
	"time"
)

const (
	apiDatelayout = "2006-01-02T15:04:05.000Z"
)

func GetNowString() string {
	return GetNow().Format(apiDatelayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
