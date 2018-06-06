// Package gen contains all the components for code generation.
package gen

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

// Start{{.ContractTypeName}}Watchers starts up the event watchers for {{.ContractTypeName}}
func Start{{.ContractTypeName}}Watchers(client bind.ContractBackend, contractAddress common.Address, eventRecvChan chan model.CivilEvent) ([]event.Subscription, error) {
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(contractAddress, client)
	if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
		return nil, err
	}

    var sub event.Subscription
	subs := []event.Subscription{}
{{if .CrawlerEvents -}}
{{- range .CrawlerEvents}}

    sub, err = startWatch{{.CrawlerMethod}}(eventRecvChan, contract)
	if err != nil {
        return nil, fmt.Errorf("Error starting start{{.WatchMethod}}: err: %v", err)
	}
	subs = append(subs, sub)

{{- end}}
{{- end}}

    return subs, nil
}

{{if .CrawlerEvents -}}
{{- range .CrawlerEvents}}

func startWatch{{.CrawlerMethod}}(eventRecvChan chan model.CivilEvent, _contract *{{$.ContractTypePackage}}.{{$.ContractTypeName}}) (event.Subscription, error) {
	opts := &bind.WatchOpts{}
    recvChan := make(chan *{{$.ContractTypePackage}}.{{.EventType}})
	sub, err := _contract.Watch{{.CrawlerMethod}}(
		opts,
		recvChan,
	{{- if .ParamValues -}}
	{{range .ParamValues}}
        []{{.Type}}{},
	{{- end}}
	{{end}}
	)
	if err != nil {
		log.Errorf("Error starting Watch{{.CrawlerMethod}}: %v", err)
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-recvChan:
				civilEvent := model.NewCivilEvent("{{.CrawlerEventName}}", event)
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
