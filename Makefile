.PHONY: build clean deploy

build:
	go get ./...
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/main ./src/main.go

deploy: build
	sls deploy --stage dev