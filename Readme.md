# Reserve Trust Fibonacci
______

## Instructions 
Expose a Fibonacci sequence generator through a web API that memoizes intermediate values. The web API should expose operations to:
* fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144)
* fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120)
* clear the data store

The web API must be written in Go, and Postgres must be used as the data store for the memoized results. Please include tests for your solution, and a Readme.md describing how to build and run it.

## Running scenarios

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
BenchmarkFibonacci1-8           16299607                75.09 ns/op           16 B/op          2 allocs/op
BenchmarkFibonacci5-8             667102              1755 ns/op             624 B/op         28 allocs/op
BenchmarkFibonacci10-8            393906              3145 ns/op             968 B/op         49 allocs/op
BenchmarkFibonacci50-8            106753             11149 ns/op            3144 B/op        171 allocs/op
BenchmarkFibonacci90-8             60796             19554 ns/op            5832 B/op        292 allocs/op
BenchmarkFibonacci900-8             6302            186754 ns/op           47656 B/op       2727 allocs/op
PASS
ok      github.com/stefanmcshane/distributed-fibonacci/app/business     7.819s
```

## Notes:
* Originally implemented in int64. This will cap out at the max int available to the arch, which would be ordinal 92 for a 64-bit arch, 46 for a 32 bit-arch
    * Could be avoided by using a big.Int, or by making a generic structure for storing data which holds the bytes, and the type