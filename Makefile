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
	rm -rf testgen/

test: clean
	mkdir -p testgen
	go build -o protoc-gen-uwentity main.go
	protoc --plugin=protoc-gen-uwentity --uwentity_out=./testgen --go_out=./testgen testdata/*.proto
	go test ./...
