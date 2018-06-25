// Package main contains commands to run
package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

var (
	packageName = kingpin.Arg("package-name", "Package name for the generated files.").
		Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	err := gen.GenerateEventHandlerLists(os.Stdout, *packageName)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(2)
	}
}
