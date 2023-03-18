pwd := $(PWD)

filesystem.build:
	GOOS=wasip1 GOARCH=wasm GOROOT=$(pwd)/../.. ../../bin/go test -c -run TestFilesystem wasip1_test.go

filesystem.run:
	wasmtime run --dir=/tmp wasip1.test
