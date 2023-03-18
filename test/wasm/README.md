# WASI implementation state

Building a binary using `GOOS=wasip1` currently works.

Example:

```
GOOS=wasip1 GOARCH=wasm GOROOT=$PWD/../.. ../../bin/go test -c -run TestFilesystem wasi_test.go
```

And some of the stblib tests are passing but there is still quite a bit of work to be done:

- Fixing the remaning issues around the filesystem interaction. The different runtimes have
different behaviors with the filesystem so it's a bit tricky. I'm currently focusing on
making `fstest` pass with both Wazero and Wastime. Current issues are: absolution symlinks
are not working and `Open` a directory needs `O_DIRECTORY` for Wasmtime. Both should be relatively
quick to fix.

- Next on my list is de-duplicating `mem_js.go` and `mem_wasi.go` and create `mem_wasm.go`.
The idea is to merge both once `resetMemoryDataView` (https://github.com/Pryz/go/blob/master/src/runtime/mem_js.go#L71)
has been moved away. See: https://github.com/WebAssembly/design/issues/1296. Here the best is
probably to have some kind of var in `mem_wasm.go` and call the function from https://github.com/Pryz/go/blob/master/src/runtime/os_js.go#L103.
The move of `resetMemoryDataView` will require it's own CL.

- Fix `sigpanic` which is currently totally disabled.

- Run all the std testsuite and fix things if needed.

## Running tests under different runtimes

Install [wasmtime](https://github.com/bytecodealliance/wasmtime/releases/tag/v5.0.0)
and the wazero CLI:

```
$ go install github.com/tetratelabs/wazero/cmd/wazero@latest
```

You can then run the tests on either runtime. By default, it will use `wasmtime`:

```
$ GOOS=wasip1 GOARCH=wasm GOROOT=$PWD/../.. PATH=$PWD/../../misc/wasm/:$PATH ../../bin/go test -timeout 10s -v ../../src/bufio/example_test.go
```

Set the environment variable `RUNTIME=wazero` to use `wazero`:

```
$ RUNTIME=wazero GOOS=wasip1 GOARCH=wasm GOROOT=$PWD/../.. PATH=$PWD/../../misc/wasm/:$PATH ../../bin/go test -timeout 10s -v ../../src/bufio/example_test.go
```
