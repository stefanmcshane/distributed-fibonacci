# Reserve Trust Fibonacci
______

## Instructions 
Expose a Fibonacci sequence generator through a web API that memoizes intermediate values. The web API should expose operations to:
* fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144)
* fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120)
* clear the data store

The web API must be written in Go, and Postgres must be used as the data store for the memoized results. Please include tests for your solution, and a Readme.md describing how to build and run it.


## Benchmarks

These can be run by running `make bench` and are calculated using an in-memory caching layer

```
go test -benchmem -bench BenchmarkFibonacci github.com/stefanmcshane/distributed-fibonacci/app/business
goos: darwin
goarch: amd64
pkg: github.com/stefanmcshane/distributed-fibonacci/app/business
cpu: Intel(R) Core(TM) i7-7920HQ CPU @ 3.10GHz
BenchmarkFibonacci1-8           433370692                2.774 ns/op           0 B/op          0 allocs/op
BenchmarkFibonacci5-8           109691816                9.260 ns/op           0 B/op          0 allocs/op
BenchmarkFibonacci10-8          81727818                13.63 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci50-8          79095604                16.57 ns/op            0 B/op          0 allocs/op
BenchmarkFibonacci90-8          74445926                14.45 ns/op            0 B/op          0 allocs/op
PASS
```

## Notes:
* This will cap out at the max int available to the arch, which would be ordinal 92 for a 64-bit arch, 46 for a 32 bit-arch
    * Could be avoided by using a big.Int, or by making a generic structure for storing data which holds the bytes, and the type