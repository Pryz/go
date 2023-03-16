// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package runtime

import "unsafe"

// https://webassembly.github.io/spec/core/exec/runtime.html#memory-instances
const _PAGESIZE = 64 * 1024

func sbrk(n uintptr) unsafe.Pointer {
	grow := (int32(n) + _PAGESIZE - 1) / _PAGESIZE
	size := currentMemory()

	if growMemory(grow) < 0 {
		return nil
	}

	return unsafe.Pointer(uintptr(size) * _PAGESIZE)
}

// Implemented in src/runtime/sys_wasm.s
func currentMemory() int32
func growMemory(pages int32) int32
