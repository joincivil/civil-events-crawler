// Package persistence contains components to interact with the DB
package persistence // import "github.com/joincivil/civil-events-crawler/pkg/persistence"

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/civil-events-crawler/pkg/model"
	"github.com/joincivil/civil-events-crawler/pkg/persistence/postgres"
)

// NewPostgresPersister creates a new postgres persister
func NewPostgresPersister(host string, port int, user string, password string, dbname string) (*PostgresPersister, error) {
	pgdb, err := postgres.NewPostgresInterface(host, port, user, password, dbname)
	if err != nil {
		return nil, err
	}
	return &PostgresPersister{
		eventToLastBlockNumber: make(map[common.Address]map[string]PersisterBlockData),
		db: pgdb}, nil
}

// PostgresPersister holds DB connection
type PostgresPersister struct {
	eventToLastBlockNumber map[common.Address]map[string]PersisterBlockData
	db                     *postgres.Interface
}

// LastBlockNumber returns the last block number seen by the persister
func (p *PostgresPersister) LastBlockNumber(eventType string, contractAddress common.Address) uint64 {
	if p.eventToLastBlockNumber[contractAddress] == nil {
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
	}
	return p.eventToLastBlockNumber[contractAddress][eventType].BlockNumber
}

// LastBlockHash returns the last block hash seen by the persister
func (p *PostgresPersister) LastBlockHash(eventType string, contractAddress common.Address) common.Hash {
	if p.eventToLastBlockNumber[contractAddress] == nil {
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
	}
	return p.eventToLastBlockNumber[contractAddress][eventType].BlockHash
}

// UpdateLastBlockData updates the last block number seen by the retriever
func (p *PostgresPersister) UpdateLastBlockData(events []model.CivilEvent) error {
	for _, event := range events {
		err := p.parseEventAndPersist(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PostgresPersister) parseEventAndPersist(event model.CivilEvent) error {
	eventType := event.EventType()
	contractAddress := event.ContractAddress()
	blockNumber := event.BlockNumber()
	blockHash := event.BlockHash()
	if p.eventToLastBlockNumber[contractAddress] == nil {
		p.eventToLastBlockNumber[contractAddress] = make(map[string]PersisterBlockData)
	}
	p.eventToLastBlockNumber[contractAddress][eventType] = PersisterBlockData{blockNumber, blockHash}
	return nil
}

// SaveEvents saves events to the DB
func (p *PostgresPersister) SaveEvents(events []*model.CivilEvent) error {
	err := p.db.SaveToEventsTable(events)
	return err
}

// PersisterBlockData is the data about block stored for persistence
type PersisterBlockData struct {
	BlockNumber uint64
	BlockHash   common.Hash
}
