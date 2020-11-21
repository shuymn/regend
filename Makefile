.PHONY: build
build: generate
	go build -o ./bin/regend ./cmd/regend/main.go

.PHONY: generate
generate:
	go generate static
