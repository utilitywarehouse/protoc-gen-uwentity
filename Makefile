BUF_VERSION := v1.12.0

install:
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

generate: build
build:
	buf generate
