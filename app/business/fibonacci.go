package business

import (
	"context"
	"fmt"
	"math/big"
	"strings"
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
	Ordinal big.Int `db:"ordinal"`
	Fib     big.Int `db:"fibonacci"`
}

// Fibonacci returns the largest fibonacci, bounded by the architecture
func (rc ReserveTrustClient) ResultsUnder(ctx context.Context, fib big.Int) (int, error) {
	count := rc.Storage.ResultsUnder(ctx, fib)
	if count == -1 {
		return -1, NewError(fib, ErrorUnableToCount)
	}
	return count, nil
}

// Fibonacci returns the largest fibonacci, bounded by the architecture
func (rc ReserveTrustClient) Fibonacci(ctx context.Context, ordinal big.Int) (big.Int, error) {
	one := big.NewInt(1)
	zero := big.NewInt(0)

	if ordinal.Cmp(zero) == -1 {
		return *zero, NewError(ordinal, ErrorInvalidOrdinal)
	} else if ordinal.Cmp(zero) == 0 || ordinal.Cmp(big.NewInt(1)) == 0 {
		return ordinal, nil
	}

	// Check if value has been calculated before, if not, get highest calculated fib
	v := big.NewInt(0).SetBytes(ordinal.Bytes())
	fib, err := rc.Storage.CheckFib(ctx, *v)
	if err == nil {
		return fib, nil
	} else {
		err = nil
	}

	// Get highest 2 stored ordinals
	var fibA, fibB big.Int
	highest, err := rc.Storage.HighestOrdinalStored(ctx)
	if err != nil && !strings.Contains(err.Error(), ErrorNoHighestOrdinal) {
		return fib, err
	}

	a, err := rc.Storage.CheckFib(ctx, highest)
	if err != nil {
		fibB = *zero
		fibA = *one
	} else {
		fibA = a
		fibB = *zero
		secondHighest := zero.Sub(&highest, one)

		if secondHighest.Cmp(big.NewInt(-1)) != 0 { // if second highest not less than 0, check the cache
			fb, err := rc.Storage.CheckFib(ctx, *secondHighest)
			if err != nil {
				return fib, NewError(*secondHighest, fmt.Sprintf("secondHighest not found in cache - %s", err.Error()))
			}
			fibB = fb
		}
	}

	// Calculate new value
	var ordinalTuplesToStore []ordinalTuple
	for highest.Cmp(&ordinal) <= 0 {
		// Add to slice for batch storage
		o := ordinalTuple{Ordinal: *big.NewInt(0).SetBytes(highest.Bytes()), Fib: *big.NewInt(0).SetBytes(fibA.Bytes())}
		ordinalTuplesToStore = append(ordinalTuplesToStore, o)

		// Store fib in a
		fibB.Add(&fibB, &fibA)
		fibA, fibB = fibB, fibA
		highest.Add(&highest, one)
	}

	// Store new calculated value
	err = rc.Storage.AddFib(ctx, ordinalTuplesToStore)
	if err != nil {
		return *big.NewInt(-1), NewError(ordinal, fmt.Sprintf("%s - %s", ErrorStorageAddError, err.Error()))
	}
	return fibB, nil
}
