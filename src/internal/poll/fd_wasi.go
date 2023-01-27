// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package poll

import "syscall"

// ReadDir wraps syscall.ReadDir.
// We treat this like an ordinary system call rather than a call
// that tries to fill the buffer.
func (fd *FD) ReadDir(buf []byte, cookie syscall.Dircookie_t) (int, error) {
	if err := fd.incref(); err != nil {
		return 0, err
	}
	defer fd.decref()
	for {
		n, err := syscall.ReadDir(fd.Sysfd, buf, cookie)
		if err != nil {
			n = 0
			if err == syscall.EAGAIN && fd.pd.pollable() {
				if err = fd.pd.waitRead(fd.isFile); err == nil {
					continue
				}
			}
		}
		// Do not call eofError; caller does not expect to see io.EOF.
		return n, err
	}
}
