// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package runtime

// Stub so we can compile the code on both js and wasi, but there is no memory
// data view to reset on wasi.
func resetMemoryDataView() {}
