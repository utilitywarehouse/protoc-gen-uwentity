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
	rm -rf testgen/testdata/

test: clean
	mkdir -p testgen/testdata/
	go build -o protoc-gen-uwentity main.go
	protoc --plugin=protoc-gen-uwentity --uwentity_out=./testgen --go_out=./testgen testgen/proto/testdata/*.proto
	go test ./...
