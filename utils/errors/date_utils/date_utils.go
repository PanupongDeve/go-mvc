package date_utils

import (
	"time"
)

const (
	// YYYY-DD-MMTHH:mm:ssZ
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNowWithTimezone(timeZone string) time.Time {
	location, _ := time.LoadLocation(timeZone)
	return time.Now().In(location)
}

func GetNow() time.Time {
	return GetNowWithTimezone("UTC")
}

func GetNowStringWithTimzone(timeZone string) string {
	location, _ := time.LoadLocation(timeZone)

	now := time.Now().In(location)

	return now.Format(apiDateLayout)
}

func GetNowString() string {
	return GetNowStringWithTimzone("UTC")
}

func GetNowStringAsiaBangkok() string {
	return GetNowStringWithTimzone("Asia/Bangkok")
}
