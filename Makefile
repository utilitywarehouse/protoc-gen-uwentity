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

test-prepare:
	 - mv testgen/proto/testdata/nested_message.proto_test testgen/proto/testdata/nested_message.proto && \
 	mv testgen/proto/testdata/simple_message.proto_test testgen/proto/testdata/simple_message.proto

test: clean test-prepare
	mkdir -p testgen/testdata/
	 - go build -o protoc-gen-uwentity main.go
	 - protoc --plugin=protoc-gen-uwentity --uwentity_out=./testgen --go_out=./testgen testgen/proto/testdata/*.proto
	 - go test ./...
	mv testgen/proto/testdata/nested_message.proto testgen/proto/testdata/nested_message.proto_test && \
	mv testgen/proto/testdata/simple_message.proto testgen/proto/testdata/simple_message.proto_test