install:
	go install github.com/bufbuild/buf/cmd/buf

generate: build
build:
	buf generate

lint:
	buf lint

format:
	buf format proto -w

clean:
	rm -rf testdata/go/*

test: clean
	go build -o protoc-gen-uwentity main.go
	protoc --plugin=protoc-gen-uwentity --uwentity_out=testdata/go --go_out=testdata/go testdata/*.proto
	go test ./...
