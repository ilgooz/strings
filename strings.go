package strings

import (
	"errors"
	"time"
)

func InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ToTime(s string) (time.Time, error) {
	date := time.Time{}
	err := date.UnmarshalText([]byte(s))
	if err != nil {
		return date, errors.New("UTC format required")
	}
	return date, nil
}
