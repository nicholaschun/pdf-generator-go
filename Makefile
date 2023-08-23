.PHONY: build clean deploy

clean:
  rm -rf ./bin

build:
	go get ./...
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/main ./src/main.go

deploy: clean build
	sls deploy --stage dev