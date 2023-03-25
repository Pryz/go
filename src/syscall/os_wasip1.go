// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

//go:wasmimport wasi_snapshot_preview1 proc_exit
func __wasip1_proc_exit(code int32)

func ProcExit(code int32) {
	__wasip1_proc_exit(code)
}
