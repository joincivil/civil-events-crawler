// Package gen contains all the components for code generation.
package gen

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/joincivil/civil-events-crawler/pkg/model"
)

const (
	defaultStartBlock = 0
)

var handlerToTemplate = map[string]TemplateData{
	"watcher": {
		tmplName: "watcher.tmpl",
		tmplVar:  watcherTmpl,
	},
	"filterer": {
		tmplName: "filterer.tmpl",
		tmplVar:  filtererTmpl,
	},
	"common": {
		tmplName: "common.tmpl",
		tmplVar:  commonTmpl,
	},
}

// TemplateData is a struct to store template information
type TemplateData struct {
	tmplName string
	tmplVar  string
}

// GenerateEventHandlers will code gen the contract event handlers for a given
// ContractType. It will output the generated code to the given io.Writer.
func GenerateEventHandlers(writer io.Writer, contractType model.ContractType, packageName string,
	handlerName string) error {

	contractSpecs, ok := model.ContractTypeToSpecs.Get(contractType)
	if !ok {
		return errors.New("Invalid ContractType")
	}

	return generateEventHandlersFromABI(writer, contractSpecs.AbiStr(), packageName,
		contractSpecs.ImportPath(), contractSpecs.TypePackage(), contractSpecs.Name(), handlerName)
}

// GenerateEventHandlersFromTemplate will code gen the contract event handlers for the
// given contract data.  It will output the generated code to the given io.Writer.
// If gofmt is true, will run go formatting on the code before output.
// packageName specifies for listener or retriever code
func GenerateEventHandlersFromTemplate(writer io.Writer, tmplData *EventHandlerContractTmplData,
	gofmt bool, handlerName string) error {
	tmpData, ok := handlerToTemplate[handlerName]
	if !ok {
		return errors.New("Invalid handlerName")
	}
	return OutputTemplatedData(writer, tmpData.tmplName, tmpData.tmplVar, tmplData, gofmt)
}

// EventHandlerMethodParam represents a value to be passed into the
// method for starting up event handlers in a smart contract.
// Maps to actions in the event handler templates.
type EventHandlerMethodParam struct {
	Type string
}

// EventHandlerTmplData represents data for an individual contract event method in a
// smart contract.
// Maps to actions in the event handler templates.
type EventHandlerTmplData struct {
	EventMethod string
	EventType   string
	ParamValues []*EventHandlerMethodParam
	EventName   string
}

// EventHandlerContractTmplData represents data for a category of contract event methods.
// Maps to actions in the event handler templates.
type EventHandlerContractTmplData struct {
	PackageName         string
	AdditionalImports   []string
	ContractImportPath  string
	ContractTypePackage string
	ContractTypeName    string
	DefaultStartBlock   int
	GenTime             time.Time
	EventHandlers       []*EventHandlerTmplData
}

func generateEventHandlersFromABI(writer io.Writer, _abiStr string, packageName string,
	contractImportPath string, contractTypePackage string, contractTypeName string,
	handlerName string) error {
	_abi, err := abi.JSON(strings.NewReader(_abiStr))
	if err != nil {
		return err
	}
	eventsIndex := 0
	eventHandlers := make([]*EventHandlerTmplData, len(_abi.Events))
	additionalImports := []string{}

	// Keep the event methods sorted by name
	sortedEvents := eventsToSortedEventsSlice(_abi.Events)
	for _, event := range sortedEvents {
		params := []*EventHandlerMethodParam{}
		for _, input := range event.Inputs {
			if input.Indexed {
				importName, paramType := translateType(input.Type.String())
				val := &EventHandlerMethodParam{Type: paramType}
				if importName != "" {
					additionalImports = append(additionalImports, importName)
				}
				params = append(params, val)
			}
		}
		eventHandler := &EventHandlerTmplData{
			EventName:   event.Name,
			EventMethod: cleanEventName(event.Name),
			EventType:   eventType(contractTypeName, event.Name),
			ParamValues: params,
		}

		eventHandlers[eventsIndex] = eventHandler
		eventsIndex++
	}
	contractData := &EventHandlerContractTmplData{
		PackageName:         packageName,
		AdditionalImports:   additionalImports,
		ContractImportPath:  contractImportPath,
		ContractTypePackage: contractTypePackage,
		ContractTypeName:    contractTypeName,
		DefaultStartBlock:   defaultStartBlock,
		GenTime:             time.Now().UTC(),
		EventHandlers:       eventHandlers,
	}
	return GenerateEventHandlersFromTemplate(writer, contractData, true, handlerName)
}

func eventType(contractTypeName string, eventName string) string {
	return fmt.Sprintf(
		"%v%v",
		contractTypeName,
		cleanEventName(eventName),
	)
}

func cleanEventName(name string) string {
	return strings.Trim(name, " _")
}

// translateType inspired by bindUnnestedTypeGo in go-ethereum/accounts/abi/bind
func translateType(stringKind string) (string, string) {
	switch {
	case strings.HasPrefix(stringKind, "address"):
		return "", "common.Address"

	case strings.HasPrefix(stringKind, "bytes"):
		parts := regexp.MustCompile(`bytes([0-9]*)`).FindStringSubmatch(stringKind)
		return "", fmt.Sprintf("[%s]byte", parts[1])

	case strings.HasPrefix(stringKind, "int") || strings.HasPrefix(stringKind, "uint"):
		parts := regexp.MustCompile(`(u)?int([0-9]*)`).FindStringSubmatch(stringKind)
		switch parts[2] {
		case "8", "16", "32", "64":
			return "", fmt.Sprintf("%sint%s", parts[1], parts[2])
		}
		return "math/big", "*big.Int"

	case strings.HasPrefix(stringKind, "bool"):
		return "", "bool"

	case strings.HasPrefix(stringKind, "string"):
		return "", "string"

	default:
		return "", stringKind
	}
}

func eventsToSortedEventsSlice(eventsMap map[string]abi.Event) []abi.Event {
	sortedEvents := make([]abi.Event, len(eventsMap))
	ind := 0
	for _, val := range eventsMap {
		sortedEvents[ind] = val
		ind++
	}
	sort.Sort(abiEventNameSort(sortedEvents))
	return sortedEvents
}

type abiEventNameSort []abi.Event

func (e abiEventNameSort) Len() int {
	return len(e)
}

func (e abiEventNameSort) Less(i, j int) bool {
	return e[i].Name < e[j].Name
}

func (e abiEventNameSort) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
