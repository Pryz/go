// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi && wasm

package runtime

import (
	"unsafe"
)

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//
//go:nosplit
func sysAlloc(n uintptr, sysStat *sysMemStat) unsafe.Pointer {
	p := sysReserve(nil, n)
	sysMap(p, n, sysStat)
	return p
}

func sysUnused(v unsafe.Pointer, n uintptr) {
}

func sysUsed(v unsafe.Pointer, n, prepared uintptr) {
}

func sysHugePage(v unsafe.Pointer, n uintptr) {
}

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//
//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	// FIXME
	sysStat.add(-int64(n))
	sysFreeOS(v, n)
}

func sysFault(v unsafe.Pointer, n uintptr) {
}

var reserveEnd uintptr

func sysReserve(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	return sysReserveOS(v, n)
}

// src/runtime/sys_wasm.s#92
func currentMemory() int32

// src/runtime/sys_wasm.s#98
func growMemory(pages int32) int32

func sysMap(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	// FIXME
	//mSysStatInc(sysStat, n)
	sysStat.add(int64(n))
	sysMapOS(v, n)
}

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//
//go:nosplit
func sysAllocOS(n uintptr) unsafe.Pointer {
	p := sysReserveOS(nil, n)
	sysMapOS(p, n)
	return p
}

func sysUnusedOS(v unsafe.Pointer, n uintptr) {
}

func sysUsedOS(v unsafe.Pointer, n uintptr) {
}

func sysHugePageOS(v unsafe.Pointer, n uintptr) {
}

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//
//go:nosplit
func sysFreeOS(v unsafe.Pointer, n uintptr) {
	// FIXME
}

func sysFaultOS(v unsafe.Pointer, n uintptr) {
}

func sysReserveOS(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	// TODO(neelance): maybe unify with mem_plan9.go, depending on how https://github.com/WebAssembly/design/blob/master/FutureFeatures.md#finer-grained-control-over-memory turns out

	if v != nil {
		// The address space of WebAssembly's linear memory is contiguous,
		// so requesting specific addresses is not supported. We could use
		// a different address, but then mheap.sysAlloc discards the result
		// right away and we don't reuse chunks passed to sysFree.
		return nil
	}

	// Round up the initial reserveEnd to 64 KiB so that
	// reservations are always aligned to the page size.
	initReserveEnd := alignUp(lastmoduledatap.end, physPageSize)
	if reserveEnd < initReserveEnd {
		reserveEnd = initReserveEnd
	}
	v = unsafe.Pointer(reserveEnd)
	reserveEnd += alignUp(n, physPageSize)

	current := currentMemory()
	// reserveEnd is always at a page boundary.
	needed := int32(reserveEnd / physPageSize)
	if current < needed {
		if growMemory(needed-current) == -1 {
			return nil
		}
		// FIXME: resetMemoryDataView()
	}

	return v
}

func sysMapOS(v unsafe.Pointer, n uintptr) {
	// FIXME
}
