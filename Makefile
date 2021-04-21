bench:
	go test -benchmem -bench BenchmarkFibonacci github.com/stefanmcshane/distributed-fibonacci/app/business

test:
	go test ./... -v -count=1

db:
	docker build . -t test/psql:1 -f build/psql.Dockerfile && \
	docker run -it -p 5432:5432 test/psql:1

api:
	go run app/api/*.go

test-int:
	curl localhost:8080/fibonacci/10  && \
	curl localhost:8080/fibonacci/25  && \
	curl localhost:8080/results-under/120 && \
	curl -i localhost:8080/clear-cache/