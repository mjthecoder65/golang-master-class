package codewars

import "strconv"

const HOUR_IN_SECONDS int = 3600
const MINUTE_IN_SECONDS int = 60

func HumanReadableTime(seconds int) string {

	hours := seconds / HOUR_IN_SECONDS

	hours_str := strconv.Itoa(hours)

	if hours < 10 {
		hours_str = "0" + hours_str
	}

	seconds = seconds % HOUR_IN_SECONDS

	minutes := seconds / MINUTE_IN_SECONDS

	minutes_str := strconv.Itoa(minutes)

	if minutes < 10 {
		minutes_str = "0" + minutes_str
	}

	seconds = seconds % MINUTE_IN_SECONDS

	seconds_str := strconv.Itoa(seconds)

	if seconds < 10 {
		seconds_str = "0" + seconds_str
	}

	return hours_str + ":" + minutes_str + ":" + seconds_str
}
