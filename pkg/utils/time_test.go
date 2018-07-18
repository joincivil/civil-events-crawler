// Package time_test contains tests for the time utils
package utils_test

import (
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"testing"
	"time"
)

var EPSILON float64 = 0.999

func floatEquals(a float64, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

func TestCurrentEpochSecsInFloat(t *testing.T) {
	ts := utils.CurrentEpochSecsInFloat()
	if ts <= 0.0 {
		t.Error("Timestamp is <= 0.0, it should be greater than 0")
	}
	now := time.Now()
	if !floatEquals(ts, float64(now.Unix())) {
		t.Error("Float timestamp is not equivalent to the calculated timestamp")
	}
}

func TestCurrentEpochSecsInInt64(t *testing.T) {
	ts := utils.CurrentEpochSecsInInt64()
	if ts <= 0 {
		t.Error("Timestamp is <= 0, it should be greater than 0")
	}
	now := time.Now()
	if now.Unix() != int64(ts) {
		t.Error("Int64 timestamp is not equal to the calculated timestamp")
	}
}

func TestCurrentEpochSecsInInt(t *testing.T) {
	ts := utils.CurrentEpochSecsInInt()
	if ts <= 0 {
		t.Error("Timestamp is <= 0, it should be greater than 0")
	}
	now := time.Now()
	if int(now.Unix()) != ts {
		t.Error("Int timestamp is not equal to the calculated timestamp")
	}
}
