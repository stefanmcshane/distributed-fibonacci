package business

import (
	"context"
	"math/big"
	"testing"
)

func TestFibonacci(t *testing.T) {
	c := testNewRTClientInMemoryStorage()
	tt := []struct {
		name          string
		ordinal       big.Int
		expectError   bool
		expectedValue *big.Int
	}{
		{
			name:          "happy path - 0",
			ordinal:       *big.NewInt(0),
			expectError:   false,
			expectedValue: big.NewInt(0),
		},
		{
			name:          "happy path - 10",
			ordinal:       *big.NewInt(10),
			expectError:   false,
			expectedValue: big.NewInt(55),
		},
		{
			name:          "happy path - 70",
			ordinal:       *big.NewInt(70),
			expectError:   false,
			expectedValue: big.NewInt(190392490709135),
		},
		{
			name:          "happy path - 50",
			ordinal:       *big.NewInt(50),
			expectError:   false,
			expectedValue: big.NewInt(12586269025),
		},
		{
			name:          "happy path - 7",
			ordinal:       *big.NewInt(7),
			expectError:   false,
			expectedValue: big.NewInt(13),
		},
		{
			name:          "happy path - 5",
			ordinal:       *big.NewInt(5),
			expectError:   false,
			expectedValue: big.NewInt(5),
		},
		{
			name:          "happy path - 92",
			ordinal:       *big.NewInt(92),
			expectError:   false,
			expectedValue: big.NewInt(7540113804746346429),
		},
		{
			name:          "happy path - 99999",
			ordinal:       *big.NewInt(99999),
			expectError:   true,
			expectedValue: big.NewInt(0),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := c.Fibonacci(context.Background(), tc.ordinal)
			if (err != nil) && !tc.expectError {
				t.Errorf("simpleFib() = unexpected error for ordinal %v - %v", tc.ordinal, err)
			}
			if got.Cmp(tc.expectedValue) != 0 && tc.expectError == false {
				t.Errorf("simpleFib() = unexpected fib value for ordinal %v - %s", tc.ordinal, got.String())
				t.Log(err)
			}
		})
	}
}

func BenchmarkFibonacci1(b *testing.B)   { benchmarkFib(b, 1, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci5(b *testing.B)   { benchmarkFib(b, 5, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci10(b *testing.B)  { benchmarkFib(b, 10, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci50(b *testing.B)  { benchmarkFib(b, 50, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci90(b *testing.B)  { benchmarkFib(b, 90, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci900(b *testing.B) { benchmarkFib(b, 900, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci900_2000Cached(b *testing.B) {
	c := testNewRTClientInMemoryStorage()
	benchmarkFib(b, 900, c)
	benchmarkFib(b, 20000, c)
}
func BenchmarkFibonacci2000_900Cached(b *testing.B) {
	c := testNewRTClientInMemoryStorage()
	benchmarkFib(b, 20000, c)
	benchmarkFib(b, 900, c)
}

func benchmarkFib(b *testing.B, topOrd int64, c ReserveTrustClient) {
	for n := 0; n < b.N; n++ {
		c.Fibonacci(context.Background(), *big.NewInt(topOrd))
	}
}

func testNewRTClientInMemoryStorage() ReserveTrustClient {
	storage := NewInMemoryStorage()
	return NewRTClient(storage)
}
