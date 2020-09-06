.PHONY: run test dep build
# NAME = $(notdir $(shell pwd))

run:
	go run *.go
dep:
	go get -v && go mod tidy
build:
	go build
test:
	go test -count=1