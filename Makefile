help:
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

test:
	goenv exec go test -v ./...

build:
	GOOS=linux GOARCH=amd64 goenv exec go build -trimpath -ldflags '-s -w' -o bin/main_linux
	GOOS=darwin GOARCH=arm64 goenv exec go build -trimpath -ldflags '-s -w' -o bin/main_mac

fmt:
	goenv exec go fmt ./...

lint:
	goenv exec go vet ./...
