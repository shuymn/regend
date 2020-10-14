.PHONY: build
build: generate
	GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -o ./bin/regend ./cmd/regend/main.go

.PHONY: generate
generate:
	go generate static
