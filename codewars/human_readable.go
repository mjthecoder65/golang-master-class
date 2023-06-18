package codewars

import "fmt"

const (
	HOUR_IN_SECONDS   int = 3600
	MINUTE_IN_SECONDS int = 60
)

func HumanReadableTime(seconds int) string {
	hours := seconds / HOUR_IN_SECONDS
	minutes := (seconds % HOUR_IN_SECONDS) / MINUTE_IN_SECONDS
	seconds = seconds % MINUTE_IN_SECONDS

	return fmt.Sprintf("%.2d:%.2d:%.2d", hours, minutes, seconds)
}
