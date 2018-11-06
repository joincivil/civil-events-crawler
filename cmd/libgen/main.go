package main

// genlib takes in input ABI and generates go files with consts for the ABI JSON
// and the bytecode.  These are used to deploy libraries to ethereum without
// the bindings.

// a bunch of code here is based on abigen, as this is basically abigen without
// the cool code bindings.
// https://github.com/ethereum/go-ethereum/blob/master/cmd/abigen/main.go

import (
	// "encoding/json"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	// "strings"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common/compiler"
	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

var (
	// Matching flags from abigen for consistency
	abiFlag = flag.String("abi", "", "Path to the Ethereum contract ABI json to bind, - for STDIN")
	binFlag = flag.String("bin", "", "Path to the Ethereum contract bytecode (generate deploy method)")
	typFlag = flag.String("type", "", "Struct name for the binding (default = package name)")
	pkgFlag = flag.String("pkg", "", "Package name to generate the binding into")
	outFlag = flag.String("out", "", "Output file for the generated binding (default = stdout)")
)

func main() {
	// Parse and ensure all needed inputs are specified
	flag.Parse()

	if *abiFlag == "" {
		fmt.Printf("No contract ABI (--abi)\n")
		os.Exit(-1)
	}

	if *binFlag == "" {
		fmt.Printf("No contract BIN (--bin)\n")
		os.Exit(-1)
	}

	if *pkgFlag == "" {
		fmt.Printf("No destination package specified (--pkg)\n")
		os.Exit(-1)
	}

	abi, err := ioutil.ReadFile(*abiFlag)
	if err != nil {
		fmt.Printf("Failed to read input ABI: %v\n", err)
		os.Exit(-1)
	}

	bin, err := ioutil.ReadFile(*binFlag)
	if err != nil {
		fmt.Printf("Failed to read input bytecode: %v\n", err)
		os.Exit(-1)
	}

	theType := *typFlag
	if theType == "" {
		theType = *pkgFlag
	}

	codeBuf := &bytes.Buffer{}
	err = gen.GenerateContractABIBIN(codeBuf, string(abi), string(bin), theType,
		*pkgFlag)
	if err != nil {
		fmt.Printf("Error w generation: %v\n", err)
		os.Exit(-1)
	}

	// Either flush it out to a file or display on the standard output
	if *outFlag == "" {
		fmt.Printf("%s\n", codeBuf.String())
		return
	}

	err = ioutil.WriteFile(*outFlag, codeBuf.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed to write ABI/BIN files: %v\n", err)
		os.Exit(-1)
	}
}
