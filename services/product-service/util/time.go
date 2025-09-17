package util

import (
	"time"
)

func GetNowWithTimeZone(timezone string) (time.Time, error) {
	var (
		now = time.Now()
	)
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return now, err
	}
	return now.In(location), nil
}
