// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package syscall

type clockid uint32

const (
	clockRealtime clockid = iota
	clockMonotonic
	clockProcessCPUTimeID
	clockThreadCPUTimeID
)

type timestamp uint64

//go:wasmimport wasi_snapshot_preview1 clock_time_get
//go:noescape
func clockTimeGet(id clockid, precision timestamp, time *timestamp) Errno
