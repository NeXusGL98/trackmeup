package utils

import "time"

func IsValidDate(date string) bool {
	_, err := time.Parse(time.DateOnly, date)

	return err == nil
}
