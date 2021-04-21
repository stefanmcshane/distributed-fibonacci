package business

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
)

func TestInMemoryStorageAdd(t *testing.T) {
	ctx := context.Background()
	tt := []struct {
		name        string
		topOrdinal  int
		expectError bool
	}{
		{
			name:        "happy path - 1",
			topOrdinal:  50,
			expectError: false,
		},
		{
			name:        "happy path - 2",
			topOrdinal:  10000,
			expectError: false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := testNewRTClientInMemoryStorage()
			for o := 2; o < tc.topOrdinal; o++ {
				v := *big.NewInt(int64(o))
				_, err := c.Storage.CheckFib(ctx, v)
				if err == nil {
					t.Errorf("CheckFib1() = unexpected match checking from in-memory cache %v", v)
				}
				err = c.Storage.AddFib(ctx, []ordinalTuple{{Ordinal: v, Fib: v}})
				if err != nil {
					t.Errorf("AddFib() = unexpected error adding to in-memory cache %v - %v", v, err)
				}
				_, err = c.Storage.CheckFib(ctx, v)
				if err != nil {
					t.Errorf("CheckFib2() = expected match not in in-memory cache %v", v)
				}
			}
			randVal := int64(rand.Intn(tc.topOrdinal))
			_, err := c.Storage.CheckFib(ctx, *big.NewInt(randVal))
			if err != nil {
				t.Errorf("CheckFib3() = expected match not in in-memory cache %d", randVal)
			}
			_, err = c.Storage.CheckFib(ctx, *big.NewInt(int64(tc.topOrdinal) - 1))
			if err != nil {
				t.Errorf("CheckFib4() = expected match not in in-memory cache %d", tc.topOrdinal)
			}
			_, err = c.Storage.CheckFib(ctx, *big.NewInt(0))
			if err != nil {
				t.Errorf("CheckFib5() = expected match not in in-memory cache %d", 0)
			}
		})
	}
}

func TestInMemoryStorageCountUnder(t *testing.T) {
	ctx := context.Background()
	max := 1000
	tt := []struct {
		name          string
		fibValue      int64
		expectedValue int64
	}{
		{
			name:          "happy path - 1",
			fibValue:      120,
			expectedValue: 12,
		},
		{
			name:          "happy path - 2",
			fibValue:      5,
			expectedValue: 5,
		},
		{
			name:          "happy path - 50",
			fibValue:      12586269025,
			expectedValue: 50,
		},
	}
	c := testNewRTClientInMemoryStorage()

	for o := 0; o < max; o++ {
		v := big.NewInt(int64(o))
		_, err := c.Fibonacci(ctx, *v)
		if err != nil {
			t.Errorf("TestInMemoryStorageCountUnder() = unexpected error adding to in-memory cache %v - %v", v, err)
		}
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := c.Storage.ResultsUnder(ctx, *big.NewInt(tc.fibValue))
			if got != int(tc.expectedValue) {
				t.Errorf("TestInMemoryStorageCountUnder() = unexpected number of variables stored under %d. Got %d, expected %d", tc.fibValue, got, tc.expectedValue)
			}
		})
	}
}
func TestInMemoryStorageClear(t *testing.T) {
	ctx := context.Background()
	max := 1000
	c := testNewRTClientInMemoryStorage()

	for o := 0; o < max; o++ {
		v := big.NewInt(int64(o))
		_, err := c.Fibonacci(ctx, *v)
		if err != nil {
			t.Errorf("TestInMemoryStorageClear() = unexpected error adding to in-memory cache %v - %v", v, err)
		}
	}
	_, err := c.Storage.CheckFib(ctx, *big.NewInt(10))
	if err != nil {
		t.Errorf("TestInMemoryStorageClear() = expected match in in-memory cache %d for ordinal 0", 0)
	}

	err = c.Storage.Clear(ctx)
	if err != nil {
		t.Errorf("TestInMemoryStorageClear() = unexpected error clearinbg in-memory cache - %v", err)
	}
	lowest, err := c.Storage.CheckFib(ctx, *big.NewInt(0))
	if err == nil {
		t.Errorf("TestInMemoryStorageClear() = expected no match in in-memory cache %d - got %d", 0, lowest.Int64())
	}
	randVal := int64(rand.Intn(max))
	randGot, err := c.Storage.CheckFib(ctx, *big.NewInt(randVal))
	if err == nil {
		t.Errorf("TestInMemoryStorageClear() = expected no match in in-memory cache %d - got %d", randVal, randGot.Int64())
	}
}
