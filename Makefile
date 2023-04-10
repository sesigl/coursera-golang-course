.PHONY: install-deps tidy build lint clean test deploy deploy-nonprod deploy-prod generate-client generate-server

install-deps:
	go get ./...

tidy:
	go mod tidy

clean:
	rm -rf ./bin ./vendor Gopkg.lock

lint:
	golangci-lint run ./src/...

build: clean test
	env go build -o bin/hello_world src/internal/task-01/hello_world.go

test:
	go test ./...
