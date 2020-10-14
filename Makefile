.PHONY: build
build: generate
	GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -o ./bin/regen ./cmd/regen/main.go

.PHONY: generate
generate:
	go generate static
