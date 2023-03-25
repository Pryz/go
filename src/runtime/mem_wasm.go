// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

package runtime

import "unsafe"

func sbrk(n uintptr) unsafe.Pointer {
	grow := divRoundUp(n, physPageSize)
	size := currentMemory()

	if growMemory(int32(grow)) < 0 {
		return nil
	}

	resetMemoryDataView()
	return unsafe.Pointer(uintptr(size) * physPageSize)
}

// Implemented in src/runtime/sys_wasm.s
func currentMemory() int32
func growMemory(pages int32) int32
