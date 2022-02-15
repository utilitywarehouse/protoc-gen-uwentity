BUF_VERSION := v1.0.0-rc12

install:
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

build:
	buf generate
