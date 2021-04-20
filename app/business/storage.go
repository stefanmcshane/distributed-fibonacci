package business

import (
	"context"
	"math/big"
)

// MemoStore is used for dealing with various storage backends, and mocking
type MemoStore interface {

	// AddFib adds a list of calculated fib numbers, allowing for batches to be added in single statement
	AddFib(ctx context.Context, tup []ordinalTuple) error

	// CheckFib returns either the fib number of the given ordinal, or the highest fib number available, lower than the ordinal
	// exactMatch is true when the correct fib is returned
	CheckFib(ctx context.Context, ordinal big.Int) (fib big.Int, exactMatch bool)

	// Len returns the total number of objects in the datastore at a give time
	Len() int

	// ResultsUnder returns the number of results under a given fibonacci value .i.e 12 results under 120
	ResultsUnder(ctx context.Context, fibValue big.Int) int

	// Clear clears the underlying cache
	Clear(ctx context.Context) error
}

type InMemoryStore struct {
	Store          map[string]big.Int
	HighestOrdinal big.Int
}

func NewInMemoryStorage() InMemoryStore {
	return InMemoryStore{
		Store: make(map[string]big.Int),
	}
}

func (s InMemoryStore) AddFib(ctx context.Context, tup []ordinalTuple) error {
	for _, t := range tup {
		// Check if current fib is higher than current largest stored
		if t.Fib.Cmp(&s.HighestOrdinal) == 1 {
			s.HighestOrdinal = t.Fib
		}
		s.Store[t.Ordinal.String()] = t.Fib
	}

	return nil
}

func (s InMemoryStore) CheckFib(ctx context.Context, ordinal big.Int) (big.Int, bool) {
	i := ordinal.String()
	if val, ok := s.Store[i]; ok {
		return val, true
	}
	return s.HighestOrdinal, false
}

func (s InMemoryStore) Len() int {
	c := 0
	for range s.Store {
		c++
	}
	return c
}

func (s InMemoryStore) ResultsUnder(ctx context.Context, fibValue big.Int) int {
	c := -1
	for _, v := range s.Store {
		if v.Cmp(&fibValue) == -1 {
			c++
		}
	}
	return c
}

func (s InMemoryStore) Clear(ctx context.Context) error {
	for k := range s.Store {
		delete(s.Store, k)
	}
	return nil
}
