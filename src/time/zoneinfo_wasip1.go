// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1 && wasm

package time

// NOTE(Xe): This value doesn't make sense for WASI as it's impossible to assume what time zone the environment is using in wasip1. The filesystem context is also undefined, so we can't really assume that any particular files are there either.
var platformZoneSources = []string{}

func initLocal() {
	// NOTE(Xe): Like the above comment, we can't assume what timezone the runtime is using.
	// All of the WASM runtimes that I've tried (wasmtime, wasmer, wazero) all feed this in UTC, so it's probably safe enough to assume this.
	// wasip2 fixes this by offering a way to get the runtime's current time zone in addition to the current time.
	localLoc.name = "UTC"
}
