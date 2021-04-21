package business

import (
	"context"
	"math/big"
)

// MemoStore is used for dealing with various storage backends, and mocking
type MemoStore interface {

	// AddFib adds a list of calculated fib numbers, allowing for batches to be added in single statement
	AddFib(ctx context.Context, tup []ordinalTuple) error

	// CheckFib returns either the fib number of the given ordinal, along with a bool of if the match was correct
	CheckFib(ctx context.Context, ordinal big.Int) (big.Int, error)

	// HighestOrdinalStored returns the highest stored ordinal. Errors include if no values are stored
	HighestOrdinalStored(ctx context.Context) (big.Int, error)

	// Len returns the total number of objects in the datastore at a give time
	Len() int

	// ResultsUnder returns the number of results under a given fibonacci value .i.e 12 results under 120
	ResultsUnder(ctx context.Context, fibValue big.Int) int

	// Clear clears the underlying cache
	Clear(ctx context.Context) error
}
