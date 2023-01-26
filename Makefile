install:
	go install github.com/bufbuild/buf/cmd/buf

generate: build
build:
	buf generate
