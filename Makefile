bench:
	go test -benchmem -bench BenchmarkFibonacci github.com/stefanmcshane/distributed-fibonacci/app/business

test:
	go test ./... -v -count=1