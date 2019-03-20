package utils

import (
	"math/big"

	log "github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

// IsInterfaceEqual returns true is the two given interfaces are equal
func IsInterfaceEqual(v1, v2 interface{}) (bool, error) {
	var err error

	switch val := v1.(type) {
	case common.Address:
		val2, ok := v2.(common.Address)
		if ok && val.Hex() == val2.Hex() {
			return true, nil
		}

	case *big.Int:
		val2, ok := v2.(*big.Int)
		if ok && val.Cmp(val2) == 0 {
			return true, nil
		}

	case *big.Float:
		val2, ok := v2.(*big.Float)
		if ok && val.Cmp(val2) == 0 {
			return true, nil
		}

	case [32]byte:
		val2, ok := v2.([32]byte)
		if ok && val == val2 {
			return true, nil
		}

	case string:
		val2, ok := v2.(string)
		if ok && val == val2 {
			return true, nil
		}

	case float64:
		val2, ok := v2.(float64)
		if ok && val == val2 {
			return true, nil
		}

	case int64:
		val2, ok := v2.(int64)
		if ok && val == val2 {
			return true, nil
		}

	case int:
		val2, ok := v2.(int)
		if ok && val == val2 {
			return true, nil
		}

	default:
		log.Infof("Unknown type: %T, %T", v1, v2)
		// Last ditch equality
		if v1 == v2 {
			return true, nil
		}
		return false, errors.Errorf("Unhandled type: %T", val)

	}

	// Last

	return false, err
}
