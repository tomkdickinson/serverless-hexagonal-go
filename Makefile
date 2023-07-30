.PHONY: build clean deploy gomodgen wire

wire:
	wire ./...

build:  clean wire
	export GO111MODULE=on
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-entry github.com/tomkdickinson/serverless-hexagonal-go/cmd/get-entry
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/list-entries github.com/tomkdickinson/serverless-hexagonal-go/cmd/list-entries

clean:
	rm -rf ./bin/**

deploy: 
	sls deploy --verbose

remove:
	sls remove --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh $(MODULE)
