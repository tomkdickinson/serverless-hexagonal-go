.PHONY: build clean deploy gomodgen wire

wire:
	wire ./...

build: gomodgen wire
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-entry github.com/tomkdickinson/serverless-go-template/cmd/get-entry
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/list-entries github.com/tomkdickinson/serverless-go-template/cmd/list-entries

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
