// Package time_test contains tests for the eth utils
package eth_test

import (
	"testing"

	"github.com/joincivil/civil-events-crawler/pkg/eth"
)

func TestNormalizeEthAddress(t *testing.T) {
	addr1 := "0x39eD84CE90Bc48DD76C4760DD0F90997Ba274F9d"
	addr2 := "0x39ed84ce90bc48dd76c4760dd0f90997ba274f9d"

	normalized1 := eth.NormalizeEthAddress(addr1)
	normalized2 := eth.NormalizeEthAddress(addr2)

	if normalized1 == "" {
		t.Errorf("Should have converted address correctly")
	}
	if normalized2 == "" {
		t.Errorf("Should have converted address correctly")
	}
	if normalized1 != normalized2 {
		t.Errorf("Addresses should have matched")
	}
}
