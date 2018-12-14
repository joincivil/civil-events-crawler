// Package time contains various common time separate by utility types
package time

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

// CurrentEpochSecsInInt64 returns the current time as a timestamp
// from epoch as type int64 in seconds.
func CurrentEpochSecsInInt64() int64 {
	return time.Now().Unix()
}

// CurrentEpochSecsInInt returns the current time as a timestamp
// from epoch as type int in seconds.
func CurrentEpochSecsInInt() int {
	return int(CurrentEpochSecsInInt64())
}

// CurrentEpochNanoSecsInInt64 returns the current time as a timestamp
// from epoch as type int64 in nanoseconds.
func CurrentEpochNanoSecsInInt64() int64 {
	return time.Now().UnixNano()
}

// SecsToNanoSecsInInt64 converts a value from secs to nanoseconds.
func SecsToNanoSecsInInt64(secs int64) int64 {
	return secs * int64(1000000000)
}

// SecsFromEpochToTime converts an int64 of seconds from epoch to Time struct
func SecsFromEpochToTime(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// NanoSecsFromEpochToTime converts an int64 of nanoseconds from epoch to Time struct
func NanoSecsFromEpochToTime(ts int64) time.Time {
	return time.Unix(0, ts)
}

// ToSecsFromEpoch converts a time.Time struct to nanoseconds from epoch.
func ToSecsFromEpoch(t *time.Time) int64 {
	return t.Unix()
}

// ToNanoSecsFromEpoch converts a time.Time struct to nanoseconds from epoch.
func ToNanoSecsFromEpoch(t *time.Time) int64 {
	return t.UnixNano()
}
