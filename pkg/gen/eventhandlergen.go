// Package gen contains all the components for code generation.
package gen

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"go/format"
	"io"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/joincivil/civil-events-crawler/pkg/generated/contract"
)

const (
	// CivilTcrContractType is the enum value for the Civil TCR type
	CivilTcrContractType ContractType = 0

	// NewsroomContractType is the enum value for the Newsroom type
	NewsroomContractType ContractType = 1
)

var (
	// NameToContractTypes is the map from a readable name to ContractType
	NameToContractTypes = NameToContractType{
		"civiltcr": CivilTcrContractType,
		"newsroom": NewsroomContractType,
	}
)

var handlerToTemplate = map[string]TemplateData{
	"watcher": TemplateData{
		tmplName: "watcher.tmpl",
		tmplVar:  watcherTmpl,
	},
	"filterer": TemplateData{
		tmplName: "filterer.tmpl",
		tmplVar:  filtererTmpl,
	},
}

// ContractType is an enum for the Civil contract type
type ContractType int

// NameToContractType is a map type of readable name to a ContractType enum value
type NameToContractType map[string]ContractType

// TemplateData is a struct to store template information
type TemplateData struct {
	tmplName string
	tmplVar  string
}

// Names returns a list of the names in NameToContractType
func (n NameToContractType) Names() []string {
	keys := make([]string, len(n))
	keyIndex := 0
	for k := range n {
		keys[keyIndex] = k
		keyIndex++
	}
	return keys
}

// GenerateCivilEventHandlers will code gen the contract event handlers for a given
// ContractType. It will output the generated code to the given io.Writer.
// Currently supports only the CivilTCR and Newsroom ContractTypes.
func GenerateCivilEventHandlers(writer io.Writer, contractType ContractType, packageName string,
	handlerName string) error {
	var err error
	switch contractType {
	case CivilTcrContractType:
		err = generateCivilTCREventHandlers(writer, packageName, handlerName)
	case NewsroomContractType:
		err = generateNewsroomEventHandlers(writer, packageName, handlerName)
	default:
		return errors.New("Invalid ContractType")
	}
	return err
}

// GenerateEventHandlersFromTemplate will code gen the contract event handlers for the
// given contract data.  It will output the generated code to the given io.Writer.
// If gofmt is true, will run go formatting on the code before output.
// packageName specifies for listener or retriever code
func GenerateEventHandlersFromTemplate(writer io.Writer, contractData *ContractData, gofmt bool,
	handlerName string) error {
	var t *template.Template
	tmpData, ok := handlerToTemplate[handlerName]
	if !ok {
		return errors.New("Invalid handlerName")
	}
	t = template.Must(template.New(tmpData.tmplName).Parse(tmpData.tmplVar))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, contractData)
	if err != nil {
		return err
	}
	output := buf.Bytes()
	if gofmt {
		output, err = format.Source(buf.Bytes())
		if err != nil {
			log.Errorf("ERROR: template generated code:\n%v", buf.String())
			return err
		}
	}
	_, err = writer.Write(output)
	return err
}

// EventHandlerMethodParam represents a value to be passed into the
// method for starting up event handlers in a Civil smart contract.
// Maps to actions in the event handler templates.
type EventHandlerMethodParam struct {
	Type string
}

// EventHandler represents data for an individual contract event method in a
// Civil smart contract.
// Maps to actions in the event handler templates.
type EventHandler struct {
	EventMethod string
	EventType   string
	ParamValues []*EventHandlerMethodParam
	EventName   string
}

// ContractData represents data for a category of contract event methods.
// Maps to actions in the event handler templates.
type ContractData struct {
	PackageName         string
	AdditionalImports   []string
	ContractImportPath  string
	ContractTypePackage string
	ContractTypeName    string
	GenTime             time.Time
	EventHandlers       []*EventHandler
}

func generateCivilTCREventHandlers(writer io.Writer, packageName string, handlerName string) error {
	contractTypePackage := "contract"
	contractImportPath := "github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	contractTypeName := "CivilTCRContract"
	abi, err := abi.JSON(strings.NewReader(contract.CivilTCRContractABI))
	if err != nil {
		return err
	}
	return generateEventHandlers(writer, abi, packageName, contractImportPath,
		contractTypePackage, contractTypeName, handlerName)
}

func generateNewsroomEventHandlers(writer io.Writer, packageName string, handlerName string) error {
	contractTypePackage := "contract"
	contractImportPath := "github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	contractTypeName := "NewsroomContract"
	abi, err := abi.JSON(strings.NewReader(contract.NewsroomContractABI))
	if err != nil {
		return err
	}
	return generateEventHandlers(writer, abi, packageName, contractImportPath,
		contractTypePackage, contractTypeName, handlerName)
}

func generateEventHandlers(writer io.Writer, _abi abi.ABI, packageName string,
	contractImportPath string, contractTypePackage string, contractTypeName string,
	handlerName string) error {
	eventsIndex := 0
	eventHandlers := make([]*EventHandler, len(_abi.Events))
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
		eventHandler := &EventHandler{
			EventName:   event.Name,
			EventMethod: cleanEventName(event.Name),
			EventType:   EventType(contractTypeName, event.Name),
			ParamValues: params,
		}

		eventHandlers[eventsIndex] = eventHandler
		eventsIndex++
	}
	contractData := &ContractData{
		PackageName:         packageName,
		AdditionalImports:   additionalImports,
		ContractImportPath:  contractImportPath,
		ContractTypePackage: contractTypePackage,
		ContractTypeName:    contractTypeName,
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
	return strings.Trim(name, "_")
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
