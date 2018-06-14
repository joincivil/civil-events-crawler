// Package eventcollector contains all the business logic for the event collection
package eventcollector // import "github.com/joincivil/civil-events-crawler/pkg/eventcollector"

// import {
//     "fmt"
//     "github.com/ethereum/go-ethereum/accounts/abi/bind"
//     "github.com/joincivil/civil-events-crawler/pkg/model"
//     "github.com/joincivil/civil-events-crawler/pkg/retriever"
//     "github.com/ethereum/go-ethereum/common"
// }

// // startEventCollection contains logic to run retriever and listener.
// func startEventCollection(client bind.ContractBackend, filterers []model.ContractFilterers,
//     watchers []model.ContractWatchers, retrieverPersister model.RetrieverMetaDataPersister,
//     listenerPersister model.ListenerMetaDataPersister) bool {

//     // map contract address to contract name and get the block number

//     lastBlockNo := retrieverPersister.LastBlockNumber(eventType, contractAddress)

//     retriever := retriever.NewCivilEventRetriever(client, lastBlockNo, )
// }
