# Reserve Trust Fibonacci
______

## Instructions 
Expose a Fibonacci sequence generator through a web API that memoizes intermediate values. The web API should expose operations to:
* fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144)
* fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120)
* clear the data store

The web API must be written in Go, and Postgres must be used as the data store for the memoized results. Please include tests for your solution, and a Readme.md describing how to build and run it.

## Running the service end to end

Run the following commands in order. This will setup a postgres instance in docker, startup the api, and send sample curl requests at the api for each of the required endpoints. Each command will need to be run in a separate shell session, that supports Make.

```
make db
```
```
make api
```
```
make test-int
```

## Running test scenarios

* Tests
```
make test
```

* Benchmarks
```
make bench
```


## Benchmarks

These can be run by running `make bench` and are calculated using an in-memory caching layer. Done on a previous commit with an int64 implementation that is limited to the size of the processor arch. See notes

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

The implementation with big.Int allows for larger fibonacci numbers to be calculated, with a performance hit. This would not be possible with int64.
This requires both more space, and more time/is 9x slower.

```
go test -benchmem -bench BenchmarkFibonacci github.com/stefanmcshane/distributed-fibonacci/app/business
goos: darwin
goarch: amd64
pkg: github.com/stefanmcshane/distributed-fibonacci/app/business
cpu: Intel(R) Core(TM) i7-7920HQ CPU @ 3.10GHz
BenchmarkFibonacci1-8                   12167211                94.12 ns/op           24 B/op          3 allocs/op
BenchmarkFibonacci5-8                    3287337               355.0 ns/op            56 B/op          7 allocs/op
BenchmarkFibonacci10-8                   3818616               416.8 ns/op            56 B/op          7 allocs/op
BenchmarkFibonacci50-8                   3496898               307.6 ns/op            56 B/op          7 allocs/op
BenchmarkFibonacci90-8                   3770887               321.1 ns/op            56 B/op          7 allocs/op
BenchmarkFibonacci900-8                  3573072               320.3 ns/op            56 B/op          7 allocs/op
BenchmarkFibonacci900_2000Cached-8       1615108               671.3 ns/op           142 B/op         14 allocs/op
BenchmarkFibonacci2000_900Cached-8       1801986               627.0 ns/op           139 B/op         14 allocs/op
PASS
```

## Notes:
* Originally implemented in int64. This will cap out at the max int available to the arch, which would be ordinal 92 for a 64-bit arch, 46 for a 32 bit-arch
    * Could be avoided by using a big.Int, or by making a generic structure for storing data which holds the bytes, and the type
* No integration tests for spinning up docker and database, no mocks written for this either as the inmemory database serves a similiar purpose as the mocks