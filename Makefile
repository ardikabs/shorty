.PHONY: build test pretty mod image-build build-linux build-windows build-darwin

OUTDIR := deploy/_output

test:
		go test -coverprofile=cover.out ./...

mod:
		go mod download

build: build-linux build-windows build-darwin

build-linux:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o deploy/_output/cli/shorty_linux_amd64 app/cli/main.go

build-windows:
		GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o deploy/_output/cli/shorty_windows_amd64 app/cli/main.go

build-darwin:
		GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o deploy/_output/cli/shorty_darwin_amd64 app/cli/main.go


pretty:
		gofmt -s -w **/*.go
		go tool vet .

image-build:
		docker build -t shorty:0.0.1 .