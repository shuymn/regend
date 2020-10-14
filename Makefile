.PHONY: build
build:
	GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -o ./bin/regen ./main.go
