package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

var (
	contractTypeName = kingpin.Arg("contract-name", "Contract watcher type to generate.").
				Required().HintAction(gen.NameToContractTypes.Names).Enum(gen.NameToContractTypes.Names()...)
	packageName = kingpin.Arg("package-name", "Package name for the generated files.").
			Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	contractType := gen.NameToContractTypes[*contractTypeName]
	err := gen.GenerateCivilWatchers(os.Stdout, contractType, *packageName)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
