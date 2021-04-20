package business

import (
	"context"
	"testing"
)

func TestFibonacci(t *testing.T) {
	c := testNewRTClientInMemoryStorage()
	tt := []struct {
		name          string
		ordinal       int64
		expectError   bool
		expectedValue int64
	}{
		{
			name:          "happy path - 0",
			ordinal:       0,
			expectError:   false,
			expectedValue: 0,
		},
		{
			name:          "happy path - 10",
			ordinal:       10,
			expectError:   false,
			expectedValue: 55,
		},
		{
			name:          "happy path - 90",
			ordinal:       90,
			expectError:   false,
			expectedValue: 2880067194370816120,
		},
		{
			name:          "happy path - 92",
			ordinal:       92,
			expectError:   false,
			expectedValue: 7540113804746346429,
		},
		{
			name:          "sad path - overflow - 93",
			ordinal:       93,
			expectError:   true,
			expectedValue: -1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := c.Fibonacci(context.Background(), tc.ordinal)
			if (err != nil) != tc.expectError {
				t.Errorf("simpleFib() = unexpected error for ordinal %d - %v", tc.ordinal, err)
			}
			if got != tc.expectedValue && tc.expectError == false {
				t.Errorf("simpleFib() = unexpected fib value for ordinal %d - %d", tc.ordinal, got)
			}
		})
	}
}

func BenchmarkFibonacci1(b *testing.B)  { benchmarkFib(b, 1, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci5(b *testing.B)  { benchmarkFib(b, 5, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci10(b *testing.B) { benchmarkFib(b, 10, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci50(b *testing.B) { benchmarkFib(b, 50, testNewRTClientInMemoryStorage()) }
func BenchmarkFibonacci90(b *testing.B) { benchmarkFib(b, 90, testNewRTClientInMemoryStorage()) }

func benchmarkFib(b *testing.B, topOrd int64, c ReserveTrustClient) {
	for n := 0; n < b.N; n++ {
		c.Fibonacci(context.Background(), topOrd)
	}
}

func testNewRTClientInMemoryStorage() ReserveTrustClient {
	storage := NewInMemoryStorage()
	return NewRTClient(storage)
}
