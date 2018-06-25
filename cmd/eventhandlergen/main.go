// Package main contains commands to run
package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/joincivil/civil-events-crawler/pkg/gen"
	"github.com/joincivil/civil-events-crawler/pkg/model"
)

var (
	contractTypeName = kingpin.Arg("contract-name", "Contract watcher type to generate.").
				Required().HintAction(model.NameToContractTypes.Names).Enum(model.NameToContractTypes.Names()...)
	packageName = kingpin.Arg("package-name", "Package name for the generated files.").
			Required().String()
	handlerName = kingpin.Arg("handler-name", "Handler name retriever, listener, or eventdef.").
			Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	contractType, ok := model.NameToContractTypes.Get(*contractTypeName)
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
