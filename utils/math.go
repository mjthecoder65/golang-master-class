package utils

import (
	"math/rand"
	"time"
)

// Generate random integer between min and max.

func RandomInt(min int, max int) int {
	if min > max {
		panic("min must be less than max")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
