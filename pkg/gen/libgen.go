// Package gen contains all the components for code generation.
package gen

import (
	"io"
	"strings"
	"time"
	"unicode"
)

// LibGenTmplData represents the data passed to the template
type LibGenTmplData struct {
	PackageName      string
	GenTime          time.Time
	ContractTypeName string
	AbiString        string
	BinString        string
}

// GenerateContractABIBIN will write the new abi/bin const file to the io.Writer
func GenerateContractABIBIN(writer io.Writer, abi string, bin string, theType string,
	packageName string) error {
	// Strip any whitespace, return chars
	abi = strings.Map(
		func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			return r
		},
		abi,
	)
	// Escape the double quotes
	abi = strings.Replace(abi, "\"", "\\\"", -1)
	bin = strings.TrimSpace(bin)

	tmplData := &LibGenTmplData{
		PackageName:      packageName,
		GenTime:          time.Now().UTC(),
		ContractTypeName: theType,
		AbiString:        abi,
		BinString:        bin,
	}

	return generate(writer, "libgen.tmpl", libgenTmpl, tmplData, true)
}
