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
	"github.com/davecgh/go-spew/spew"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joincivil/civil-events-crawler/pkg/model"

	ctime "github.com/joincivil/go-common/pkg/time"
{{if .ContractImportPath -}}
	"{{.ContractImportPath}}"
{{- end}}
{{if .AdditionalImports -}}
{{- range .AdditionalImports}}
	"{{.}}"
{{- end}}
{{- end}}
)

func New{{.ContractTypeName}}Watchers(contractAddress common.Address) *{{.ContractTypeName}}Watchers {
	return &{{.ContractTypeName}}Watchers{
		contractAddress: contractAddress,
	}
}

type {{.ContractTypeName}}Watchers struct {
	contractAddress common.Address
	contract *{{.ContractTypePackage}}.{{.ContractTypeName}}
	activeSubs []event.Subscription
}

func (w *{{.ContractTypeName}}Watchers) ContractAddress() common.Address {
	return w.contractAddress
}

func (w *{{.ContractTypeName}}Watchers) ContractName() string {
	return "{{.ContractTypeName}}"
}

func (w *{{.ContractTypeName}}Watchers) StopWatchers() error {
	for _, sub := range w.activeSubs {
		sub.Unsubscribe()
	}
	w.activeSubs = nil
	return nil
}

func (w *{{.ContractTypeName}}Watchers) StartWatchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
	return w.Start{{.ContractTypeName}}Watchers(client, eventRecvChan)
}

// Start{{.ContractTypeName}}Watchers starts up the event watchers for {{.ContractTypeName}}
func (w *{{.ContractTypeName}}Watchers) Start{{.ContractTypeName}}Watchers(client bind.ContractBackend,
	eventRecvChan chan *model.Event) ([]event.Subscription, error) {
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(w.contractAddress, client)
	if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
		return nil, err
	}
	w.contract = contract

    var sub event.Subscription
	subs := []event.Subscription{}
{{if .EventHandlers -}}
{{- range .EventHandlers}}

    sub, err = w.startWatch{{.EventMethod}}(eventRecvChan)
	if err != nil {
        return nil, fmt.Errorf("Error starting start{{.EventMethod}}: err: %v", err)
	}
	subs = append(subs, sub)

{{- end}}
{{- end}}

	w.activeSubs = subs
    return subs, nil
}

{{if .EventHandlers -}}
{{- range .EventHandlers}}

func (w *{{$.ContractTypeName}}Watchers) startWatch{{.EventMethod}}(eventRecvChan chan *model.Event) (event.Subscription, error) {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		maxRetries := 5
		startupFn := func() (event.Subscription, chan *{{$.ContractTypePackage}}.{{.EventType}}, error) {
			retry := 0
			for {
				opts := &bind.WatchOpts{}
				recvChan := make(chan *{{$.ContractTypePackage}}.{{.EventType}})
				log.Infof("startupFn: Starting Watch{{.EventMethod}}")
				sub, err := w.contract.Watch{{.EventMethod}}(
					opts,
					recvChan,
					{{- if .ParamValues -}}
					{{range .ParamValues}}
						[]{{.Type}}{},
					{{- end}}
					{{end}}
				)
				if err != nil {
					if sub != nil {
						log.Infof("startupFn: Unsubscribing Watch{{.EventMethod}}")
						sub.Unsubscribe()
					}
					if retry >= maxRetries {
						return nil, nil, err
					}
					retry++
					log.Warningf("startupFn: Retrying start Watch{{.EventMethod}}: retry: %v: %v", retry, err)
					continue
				}
				log.Infof("startupFn: Watch{{.EventMethod}} started")
				return sub, recvChan, nil
			}
		}
		sub, recvChan, err := startupFn()
		if err != nil {
			log.Errorf("Error starting Watch{{.EventMethod}}: %v", err)
			return err
		}
		defer sub.Unsubscribe()
		log.Infof("Starting up Watch{{.EventMethod}} for contract %v", w.contractAddress.Hex())
		for {
			select {
			// 15 min premptive resubscribe
			case <-time.After(time.Second * time.Duration(60*15)):
				log.Infof("Premptive restart of {{.EventMethod}}")
				oldSub := sub
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("Error starting {{.EventMethod}}: %v", err)
					return err
				}
				oldSub.Unsubscribe()
				log.Infof("Done preemptive restart {{.EventMethod}}")
			case event := <-recvChan:
				if log.V(2) {
					log.Infof("Received event on Watch{{.EventMethod}}: %v", spew.Sprintf("%#+v", event))
				} else {
					log.Info("Received event on Watch{{.EventMethod}}")
				}
				modelEvent, err := model.NewEventFromContractEvent("{{.EventMethod}}", w.ContractName(), w.contractAddress, event, ctime.CurrentEpochSecsInInt64(), model.Watcher)
				if err != nil {
					log.Errorf("Error creating new event: event: %v, err: %v", event, err)
					continue
				}
				select {
				case eventRecvChan <- modelEvent:
					if log.V(2) {
						log.Infof("Sent event to eventRecvChan on Watch{{.EventMethod}}: %v", spew.Sprintf("%#+v", event))
					} else {
						log.Info("Sent event to eventRecvChan on Watch{{.EventMethod}}")
					}
				case err := <-sub.Err():
					log.Errorf("Error with Watch{{.EventMethod}}, fatal (a): %v", err)
					sub.Unsubscribe()
					sub, recvChan, err = startupFn()
					if err != nil {
						log.Errorf("Error restarting Watch{{.EventMethod}}, fatal (a): %v", err)
						return err
					}
					log.Errorf("Done error with Watch{{.EventMethod}}, fatal (a): %v", err)
				case <-quit:
					log.Infof("Quit Watch{{.EventMethod}} (a): %v", err)
					return nil
				}
			case err := <-sub.Err():
				log.Errorf("Error with Watch{{.EventMethod}}, fatal (b): %v", err)
				sub.Unsubscribe()
				sub, recvChan, err = startupFn()
				if err != nil {
					log.Errorf("WATCHER: Error restarting Watch{{.EventMethod}}, fatal (b): %v", err)
					return err
				}
			case <-quit:
				log.Infof("Quit Watch{{.EventMethod}} (b): %v", err)
				return nil
			}
		}
	}), nil
}

{{- end}}
{{- end}}
`
