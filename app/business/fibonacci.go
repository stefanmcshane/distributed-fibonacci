package business

import (
	"context"
	"fmt"
	"math/big"
)

type ReserveTrustClient struct {
	Storage MemoStore
}

func NewRTClient(storage MemoStore) ReserveTrustClient {
	return ReserveTrustClient{
		Storage: storage,
	}
}

type ordinalTuple struct {
	Ordinal big.Int
	Fib     big.Int
}

// Fibonacci returns the largest fibonacci, bounded by the architecture
func (rc ReserveTrustClient) Fibonacci(ctx context.Context, ordinal big.Int) (big.Int, error) {
	if ordinal.Cmp(big.NewInt(0)) == -1 {
		return *big.NewInt(0), NewError(ordinal, ErrorInvalidOrdinal)
	} else if ordinal.Cmp(big.NewInt(0)) == 0 || ordinal.Cmp(big.NewInt(1)) == 0 {
		return ordinal, nil
	}

	// Check if value has been calculated before, if not, get highest calculated fib
	v := big.NewInt(0).SetBytes(ordinal.Bytes())
	v = v.Add(v, big.NewInt(1))
	fib, exactMatch := rc.Storage.CheckFib(ctx, *v)
	if exactMatch {
		return fib, nil
	}

	// Calculate new value
	var ordinalTuplesToStore []ordinalTuple
	ordinalTuplesToStore = append(ordinalTuplesToStore, ordinalTuple{Ordinal: *big.NewInt(0), Fib: *big.NewInt(0)})
	ordinalTuplesToStore = append(ordinalTuplesToStore, ordinalTuple{Ordinal: *big.NewInt(1), Fib: *big.NewInt(1)})

	one := big.NewInt(1)
	a := big.NewInt(0)
	b := big.NewInt(1)

	var lowestFib *big.Int
	if fib.Cmp(big.NewInt(0)) == 0 {
		lowestFib = big.NewInt(1)
	} else {
		lowestFib = &fib
	}

	for lowestFib.Cmp(&ordinal) <= 0 {
		// Add to slice for batch storage
		o := ordinalTuple{Ordinal: *big.NewInt(0).SetBytes(lowestFib.Bytes()), Fib: *big.NewInt(0).SetBytes(a.Bytes())}
		ordinalTuplesToStore = append(ordinalTuplesToStore, o)

		// Store fib in a
		a.Add(a, b)
		a, b = b, a
		lowestFib.Add(lowestFib, one)
	}

	// Store new calculated value
	err := rc.Storage.AddFib(ctx, ordinalTuplesToStore)
	if err != nil {
		return *big.NewInt(-1), NewError(ordinal, fmt.Sprintf("%s - %s", ErrorStorageAddError, err.Error()))
	}

	return *a, nil
}
