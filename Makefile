.PHONY: build test pretty mod

OUTDIR := deploy/_output

test:
		go test -coverprofile=cover.out ./...

mod:
		go mod download

build:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o deploy/_output/cli/shorty app/cli/main.go

pretty:
		gofmt -s -w **/*.go
		go tool vet .