package persistence

import (
	"errors"
)

var (
	// ErrPersisterNoResults is an error that represents no results from
	// the persister on queries.  Should be returned by the persisters
	// on event of no results in retrieval queries
	ErrPersisterNoResults = errors.New("No results from persister")
)
