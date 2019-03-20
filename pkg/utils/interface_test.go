package utils_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/civil-events-crawler/pkg/utils"
)

func TestIsInterfaceDifferingTypes(t *testing.T) {
	var var1 interface{}
	var var2 interface{}

	var1 = common.HexToAddress("0x3E39fa983abcD349d95aEd608e798817397cF0D1")
	var2 = big.NewInt(1000)

	eq, err := utils.IsInterfaceEqual(var1, var2)
	if err != nil {
		t.Errorf("Should have not returned an error: err: %v", err)
	}
	if eq {
		t.Errorf("Should have returned not been equal for different type")
	}
}

func TestIsInterfaceEqualAddress(t *testing.T) {
	var var1 interface{}
	var var2 interface{}

	var1 = common.HexToAddress("0x3E39fa983abcD349d95aEd608e798817397cF0D1")
	var2 = common.HexToAddress("0x3E39fa983abcD349d95aEd608e798817397cF0D1")

	eq, err := utils.IsInterfaceEqual(var1, var2)
	if err != nil {
		t.Errorf("Should have not returned an error: err: %v", err)
	}
	if !eq {
		t.Errorf("Should have returned equals for common.Address")
	}

	var1 = common.HexToAddress("0x3E39fa983abcD349d95aEd608e798817397cF0D1")
	var2 = common.HexToAddress("0x3E38fa983abcD349d95aEd608e798817397cF0D1")

	eq, err = utils.IsInterfaceEqual(var1, var2)
	if err != nil {
		t.Errorf("Should have not returned an error: err: %v", err)
	}
	if eq {
		t.Errorf("Should have returned not equal for common.Address")
	}
}

func TestIsInterfaceEqualBigInt(t *testing.T) {
	var var1 interface{}
	var var2 interface{}

	var1 = big.NewInt(1000)
	var2 = big.NewInt(1000)

	eq, err := utils.IsInterfaceEqual(var1, var2)
	if err != nil {
		t.Errorf("Should have not returned an error: err: %v", err)
	}
	if !eq {
		t.Errorf("Should have returned equals for *big.Int")
	}

	var1 = big.NewInt(1000)
	var2 = big.NewInt(5000)

	eq, err = utils.IsInterfaceEqual(var1, var2)
	if err != nil {
		t.Errorf("Should have not returned an error: err: %v", err)
	}
	if eq {
		t.Errorf("Should have returned not equals for *big.Int")
	}
}

// func TestIsInterfaceEqual32Byte(t *testing.T) {
// 	var var1 interface{}
// 	var var2 interface{}

// 	var1 = [32]byte{}
// 	var2 = [32]byte{}

// 	eq, err := utils.IsInterfaceEqual(var1, var2)
// 	if err != nil {
// 		t.Errorf("Should have not returned an error: err: %v", err)
// 	}
// 	if !eq {
// 		t.Errorf("Should have returned equals for *big.Int")
// 	}

// 	var1 = big.NewInt(1000)
// 	var2 = big.NewInt(5000)

// 	eq, err = utils.IsInterfaceEqual(var1, var2)
// 	if err != nil {
// 		t.Errorf("Should have not returned an error: err: %v", err)
// 	}
// 	if eq {
// 		t.Errorf("Should have returned not equals for *big.Int")
// 	}
// }

type Unknown struct {
}

func TestIsInterfaceUnknownType(t *testing.T) {
	var var1 interface{}
	var var2 interface{}

	var1 = Unknown{}
	var2 = Unknown{}

	eq, err := utils.IsInterfaceEqual(var1, var2)
	if err != nil {
		t.Errorf("Should have not returned an error: err: %v", err)
	}
	if !eq {
		t.Errorf("Should have returned equals for Unknown type")
	}

}
