// Package main contains commands to run
package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	specs "github.com/joincivil/civil-events-crawler/pkg/contractspecs"
	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

var (
	contractTypeName = kingpin.Arg("contract-name", "Contract watcher type to generate.").
				Required().HintAction(specs.NameToContractTypes.Names).Enum(specs.NameToContractTypes.Names()...)
	packageName = kingpin.Arg("package-name", "Package name for the generated files.").
			Required().String()
	handlerName = kingpin.Arg("handler-name", "Handler name retriever, or listener or common").
			Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	contractType, ok := specs.NameToContractTypes.Get(*contractTypeName)
	if !ok {
		fmt.Printf("ERROR: Contract type for not found for: %v\n", *contractTypeName)
		os.Exit(2)
	}
	err := gen.GenerateEventHandlers(os.Stdout, contractType, *packageName, *handlerName)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(2)
	}

}
