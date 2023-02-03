pwd := $(PWD)

filesystem.build:
	GOOS=wasi GOARCH=wasm GOROOT=$(pwd)/../.. ../../bin/go test -c -run TestFilesystem wasi_test.go

filesystem.run:
	wasmtime run --dir=/tmp wasi.test
