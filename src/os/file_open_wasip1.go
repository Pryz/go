// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package os

import (
	"internal/poll"
	"syscall"
)

func open(path string, flag int, _ uint32) (int, poll.SysFile, error) {
	fd, path, err := syscall.PathOpen(path, flag)
	return fd, poll.SysFile{Path: path}, err
}
