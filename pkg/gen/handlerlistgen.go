// Package gen contains all the components for code generation.
package gen

import (
	"io"
	"time"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"

	cgen "github.com/joincivil/go-common/pkg/gen"
)

// EventHandlerListContractTmplData represents the names of the contracts to
// list
type EventHandlerListContractTmplData struct {
	Name       string
	SimpleName string
}

// EventHandlerListTmplData represents the data passed to the EventHandlerList
// template.
type EventHandlerListTmplData struct {
	PackageName string
	GenTime     time.Time
	Contracts   []*EventHandlerListContractTmplData
}

// GenerateEventHandlerLists will code gen a function around the event filterers/watchers
// to return a list of filterers/watchers based on a map of contract simple names to contract
// address.
func GenerateEventHandlerLists(writer io.Writer, packageName string) error {
	contracts := []*EventHandlerListContractTmplData{}
	for _, t := range specs.ContractTypeToSpecs.Types() {
		spec, _ := specs.ContractTypeToSpecs.Get(t)
		contract := &EventHandlerListContractTmplData{
			Name:       spec.Name(),
			SimpleName: spec.SimpleName(),
		}
		contracts = append(contracts, contract)
	}
	tmplData := &EventHandlerListTmplData{
		PackageName: packageName,
		Contracts:   contracts,
		GenTime:     time.Now().UTC(),
	}
	return cgen.OutputTemplatedData(writer, "handlerlist.tmpl", handlerListTmpl, tmplData, true)
}
