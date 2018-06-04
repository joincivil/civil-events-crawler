// Package utils contains various common utils separate by utility types
package utils

import (
	"time"
)

// CurrentEpochSecsInFloat returns the current time as a timestamp
// from epoch as type float64 in seconds.
func CurrentEpochSecsInFloat() float64 {
	now := time.Now()
	ts := float64(now.UnixNano()) / float64(1000*1000*1000)
	return ts
}

// CurrentEpochSecsInInt returns the current time as a timestamp
// from epoch as type float64 in seconds.
func CurrentEpochSecsInInt() int {
	return int(CurrentEpochSecsInFloat())
}
