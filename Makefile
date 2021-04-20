bench:
	go test -benchmem -bench BenchmarkFibonacci github.com/stefanmcshane/distributed-fibonacci/app/business

test:
	go test -timeout 30s github.com/stefanmcshane/distributed-fibonacci/app/business -v