package business

import (
	"context"
	"errors"
	"math/big"
)

const (
	ErrorNoHighestOrdinal = "no stored ordinals"
)

// InMemoryStore implements the MemoStore interface with an inmemory cache for testing/mocking without dependencies
type InMemoryStore struct {
	Store          map[string]big.Int
	HighestOrdinal big.Int
}

func NewInMemoryStorage() InMemoryStore {
	s := make(map[string]big.Int)
	s["0"] = *big.NewInt(0)
	s["1"] = *big.NewInt(1)
	return InMemoryStore{
		Store:          s,
		HighestOrdinal: *big.NewInt(1),
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

func (s InMemoryStore) HighestOrdinalStored(ctx context.Context) (big.Int, error) {
	if len(s.HighestOrdinal.Bits()) == 0 {
		return s.HighestOrdinal, errors.New(ErrorNoHighestOrdinal)
	}
	return s.HighestOrdinal, nil
}

func (s InMemoryStore) CheckFib(ctx context.Context, ordinal big.Int) (big.Int, error) {
	i := ordinal.String()
	if val, ok := s.Store[i]; ok {
		return val, nil
	}
	return big.Int{}, NewError(ordinal, ErrorStorageOrdinalNotFound)
}

func (s InMemoryStore) Len() int {
	c := 0
	for range s.Store {
		c++
	}
	return c
}

func (s InMemoryStore) ResultsUnder(ctx context.Context, fibValue big.Int) int {
	c := 0
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
