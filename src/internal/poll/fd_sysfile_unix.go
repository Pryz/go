// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm)

package poll

import "syscall"

type SysFile struct {
	// Writev cache.
	iovecs *[]syscall.Iovec
}

// dupCloseOnExecOld is the traditional way to dup an fd and
// set its O_CLOEXEC bit, using two system calls.
func dupCloseOnExecOld(fd int) (int, string, error) {
	syscall.ForkLock.RLock()
	defer syscall.ForkLock.RUnlock()
	newfd, err := syscall.Dup(fd)
	if err != nil {
		return -1, "dup", err
	}
	syscall.CloseOnExec(newfd)
	return newfd, "", nil
}
