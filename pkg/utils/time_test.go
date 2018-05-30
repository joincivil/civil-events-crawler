// Package time_test contains tests for the time utils
package utils_test

import (
	"github.com/joincivil/civil-events-crawler/pkg/utils"
	"testing"
)

func TestCurrentEpochSecsInFloat(t *testing.T) {
	ts := utils.CurrentEpochSecsInFloat()
	if ts <= 0.0 {
		t.Error("Timestamp is <= 0.0, it should be greater than 0")
	}
}

func TestCurrentEpochSecsInInt(t *testing.T) {
	ts := utils.CurrentEpochSecsInInt()
	if ts <= 0 {
		t.Error("Timestamp is <= 0, it should be greater than 0")
	}
}
