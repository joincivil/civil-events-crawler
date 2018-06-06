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

// ContractType is an enum for the Civil contract type
type ContractType int

// NameToContractType is a map type of readable name to a ContractType enum value
type NameToContractType map[string]ContractType

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

// GenerateCivilWatchers will code gen the contract event watchers for a given
// ContractType. It will output the generated code to the given io.Writer.
// Currently supports only the CivilTCR and Newsroom ContractTypes.
func GenerateCivilWatchers(writer io.Writer, contractType ContractType, packageName string) error {
	var err error
	switch contractType {
	case CivilTcrContractType:
		err = generateCivilTCRWatchers(writer, packageName)
	case NewsroomContractType:
		err = generateNewsroomWatchers(writer, packageName)
	default:
		return errors.New("Invalid ContractType")
	}
	return err
}

// GenerateWatchersFromTemplate will code gen the contract event watchers for the
// given Watchers data.  It will output the generated code to the given io.Writer.
// If gofmt is true, will run go formatting on the code before output.
func GenerateWatchersFromTemplate(writer io.Writer, watchers *ContractWatchers, gofmt bool) error {
	t := template.Must(template.New("watcher.tmpl").Parse(watcherTmpl))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, watchers)
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

// WatchEventMethodParam represents a value to be passed into the
// method for starting up watchers in a Civil smart contract.
// Maps to actions in the watchers.tmpl template.
type WatchEventMethodParam struct {
	Type string
}

// WatchEvent represents data for an individual contract Watch* event method in a
// Civil smart contract.
// Maps to actions in the watchers.tmpl template.
type WatchEvent struct {
	WatchMethod    string
	EventType      string
	ParamValues    []*WatchEventMethodParam
	WatchEventName string
}

// ContractWatchers represents data for a category of contract Watch* event methods.
// Maps to actions in the watchers.tmpl template.
type ContractWatchers struct {
	PackageName         string
	AdditionalImports   []string
	ContractImportPath  string
	ContractTypePackage string
	ContractTypeName    string
	GenTime             time.Time
	WatchEvents         []*WatchEvent
}

func generateCivilTCRWatchers(writer io.Writer, packageName string) error {
	contractTypePackage := "contract"
	contractImportPath := "github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	contractTypeName := "CivilTCRContract"
	abi, err := abi.JSON(strings.NewReader(contract.CivilTCRContractABI))
	if err != nil {
		return err
	}
	return generateWatchers(writer, abi, packageName, contractImportPath,
		contractTypePackage, contractTypeName)
}

func generateNewsroomWatchers(writer io.Writer, packageName string) error {
	contractTypePackage := "contract"
	contractImportPath := "github.com/joincivil/civil-events-crawler/pkg/generated/contract"
	contractTypeName := "NewsroomContract"
	abi, err := abi.JSON(strings.NewReader(contract.NewsroomContractABI))
	if err != nil {
		return err
	}
	return generateWatchers(writer, abi, packageName, contractImportPath,
		contractTypePackage, contractTypeName)
}

func generateWatchers(writer io.Writer, abi abi.ABI, packageName string,
	contractImportPath string, contractTypePackage string, contractTypeName string) error {
	eventsIndex := 0
	watchEvents := make([]*WatchEvent, len(abi.Events))
	additionalImports := []string{}

	for _, event := range abi.Events {
		params := []*WatchEventMethodParam{}
		for _, input := range event.Inputs {
			if input.Indexed {
				importName, paramType := translateType(input.Type.String())
				val := &WatchEventMethodParam{Type: paramType}
				if importName != "" {
					additionalImports = append(additionalImports, importName)
				}
				params = append(params, val)
			}
		}
		watchEvent := &WatchEvent{
			WatchEventName: event.Name,
			WatchMethod:    cleanEventName(event.Name),
			EventType:      watchEventType(contractTypeName, event.Name),
			ParamValues:    params,
		}

		watchEvents[eventsIndex] = watchEvent
		eventsIndex++
	}
	watchers := &ContractWatchers{
		PackageName:         packageName,
		AdditionalImports:   additionalImports,
		ContractImportPath:  contractImportPath,
		ContractTypePackage: contractTypePackage,
		ContractTypeName:    contractTypeName,
		GenTime:             time.Now().UTC(),
		WatchEvents:         watchEvents,
	}
	return GenerateWatchersFromTemplate(writer, watchers, true)
}

func watchEventType(contractTypeName string, eventName string) string {
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

const watcherTmpl = `
// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at {{.GenTime}}
package {{.PackageName}}

import (
	log "github.com/golang/glog"
    "fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/model"
{{if .ContractImportPath -}}
	"{{.ContractImportPath}}"
{{- end}}
{{if .AdditionalImports -}}
{{- range .AdditionalImports}}
	"{{.}}"
{{- end}}
{{- end}}
)

type {{.ContractTypeName}}Watchers struct {}

func (w *{{.ContractTypeName}}Watchers) ContractName() string {
	return "{{.ContractTypeName}}"
}

func (w *{{.ContractTypeName}}Watchers) StartWatchers(client bind.ContractBackend, contractAddress common.Address,
	eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
	return w.Start{{.ContractTypeName}}Watchers(client, contractAddress, eventRecvChan)
}

// Start{{.ContractTypeName}}Watchers starts up the event watchers for {{.ContractTypeName}}
func (w *{{.ContractTypeName}}Watchers) Start{{.ContractTypeName}}Watchers(client bind.ContractBackend,
	contractAddress common.Address, eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(contractAddress, client)
	if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
		return nil, err
	}

    var sub event.Subscription
	subs := []event.Subscription{}
{{if .WatchEvents -}}
{{- range .WatchEvents}}

    sub, err = startWatch{{.WatchMethod}}(eventRecvChan, contract)
	if err != nil {
        return nil, fmt.Errorf("Error starting start{{.WatchMethod}}: err: %v", err)
	}
	subs = append(subs, sub)

{{- end}}
{{- end}}

    return subs, nil
}

{{if .WatchEvents -}}
{{- range .WatchEvents}}

func startWatch{{.WatchMethod}}(eventRecvChan chan model.CivilEvent, _contract *{{$.ContractTypePackage}}.{{$.ContractTypeName}}) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
    recvChan := make(chan *{{$.ContractTypePackage}}.{{.EventType}})
	sub, err := _contract.Watch{{.WatchMethod}}(
		opts,
		recvChan,
	{{- if .ParamValues -}}
	{{range .ParamValues}}
        []{{.Type}}{},
	{{- end}}
	{{end}}
	)
	if err != nil {
		log.Errorf("Error starting Watch{{.WatchMethod}}: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("{{.WatchEventName}}", event)
				select {
				case eventRecvChan <- *civilEvent:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

{{- end}}
{{- end}}
`
