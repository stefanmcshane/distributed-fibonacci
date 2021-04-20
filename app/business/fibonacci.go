package business

import (
	"context"
	"fmt"
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
	Ordinal int64
	Fib     int64
}

// Fibonacci returns the largest fibonacci, bounded by the architecture
func (rc ReserveTrustClient) Fibonacci(ctx context.Context, ordinal int64) (int64, error) {
	if ordinal < 0 {
		return 0, NewError(ordinal, ErrorInvalidOrdinal)
	} else if ordinal == 0 || ordinal == 1 {
		return ordinal, nil
	}

	// Check if value has been calculated before, if not, get highest calculated fib
	fib, exactMatch := rc.Storage.CheckFib(ctx, ordinal)
	if exactMatch {
		return fib, nil
	}

	// Calculate new value
	cache := make(map[int64]int64)
	cache[0] = 0
	cache[1] = 1
	fib, ordinalCacheToStore := simpleFib(ordinal, cache)

	var ordinalTuples []ordinalTuple
	for k, v := range ordinalCacheToStore {
		ordinalTuples = append(ordinalTuples, ordinalTuple{Ordinal: k, Fib: v})
	}
	fmt.Println(ordinalTuples)
	// Store new calculated value
	err := rc.Storage.AddFib(ctx, ordinalTuples)
	if err != nil {
		return -1, NewError(ordinal, fmt.Sprintf("%s - %s", ErrorStorageAddError, err.Error()))
	}
	cache = make(map[int64]int64)

	if fib < 0 {
		return -2, NewError(ordinal, ErrorOverflow)
	}
	return fib, nil
}

// simpleFib calculates a memoized fibonacci up to the archtectural bounds, returning the fib itself, and a cache that should be stored
func simpleFib(ordinalToFind int64, tempCache map[int64]int64) (int64, map[int64]int64) {
	if val, ok := tempCache[ordinalToFind]; ok {
		return val, tempCache
	}
	v2, tempCache := simpleFib(ordinalToFind-2, tempCache)
	v1, tempCache := simpleFib(ordinalToFind-1, tempCache)
	tempCache[ordinalToFind] = v2 + v1
	return tempCache[ordinalToFind], tempCache
}
