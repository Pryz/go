// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package time

// in wasip1 zoneinfo is managed by the runtime.
var platformZoneSources = []string{}

func initLocal() {
	localLoc.name = "Local"
}
