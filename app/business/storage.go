package business

import "context"

// MemoStore is used for dealing with various storage backends, and mocking
type MemoStore interface {

	// AddFib adds a list of calculated fib numbers, allowing for batches to be added in single statement
	AddFib(ctx context.Context, tup []ordinalTuple) error

	// CheckFib returns either the fib number of the given ordinal, or the highest fib number available, lower than the ordinal
	// exactMatch is true when the correct fib is returned
	CheckFib(ctx context.Context, ordinal int64) (fib int64, exactMatch bool)
}

type InMemoryStore struct {
	Store          map[int64]int64
	HighestOrdinal int64
}

func NewInMemoryStorage() InMemoryStore {
	return InMemoryStore{
		Store: make(map[int64]int64),
	}
}

func (s InMemoryStore) AddFib(ctx context.Context, tup []ordinalTuple) error {
	for _, t := range tup {
		if t.Fib > s.HighestOrdinal {
			s.HighestOrdinal = t.Fib
		}
		s.Store[t.Ordinal] = t.Fib
	}
	return nil
}

func (s InMemoryStore) CheckFib(ctx context.Context, ordinal int64) (int64, bool) {
	if val, ok := s.Store[ordinal]; ok {
		return val, true
	}
	return s.HighestOrdinal, false
}
