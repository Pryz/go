// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package os

import (
	"io"
	"runtime"
	"syscall"
)

// Auxiliary information if the File describes a directory
type dirInfo struct {
	buf    []byte              // buffer for directory I/O
	data   []byte              // unread portion of buf
	cookie syscall.Dircookie_t // pointer to next directory entry
}

const (
	// More than 5760 to work around https://golang.org/issue/24015.
	blockSize = 8192
)

func (d *dirInfo) close() {}

func (f *File) seekInvalidate() {}

func (f *File) readdirnames(n int) (names []string, err error) {
	d := f.dirinfo
	if d == nil {
		d = &dirInfo{
			buf: make([]byte, blockSize),
		}
		f.dirinfo = d
	}

	size := n
	if size <= 0 {
		size = 100
		n = -1
	}

	names = make([]string, 0, size)
	for n != 0 {
		done := false
		if len(d.data) == 0 {
			n, err := f.pfd.ReadDir(d.buf, d.cookie)
			done = n < len(d.buf)
			d.data = d.buf[:n]
			runtime.KeepAlive(f)
			if err != nil {
				return names, wrapSyscallError("readdirnames", err)
			}
		}

		for n != 0 {
			if len(d.data) < 24 {
				d.data = nil
				break // incomplete dirent
			}
			namelen := readUint32(d.data[16:])
			if len(d.data) < int(24+namelen) {
				d.data = nil
				break // incomplete name
			}
			d.cookie = syscall.Dircookie_t(readUint64(d.data[0:]))
			name := string(d.data[24 : 24+namelen])
			d.data = d.data[24+namelen:]
			if name == "." || name == ".." {
				continue
			}
			names = append(names, name)
			n--
		}

		if done {
			break
		}
	}

	if n >= 0 && len(names) == 0 {
		return names, io.EOF
	}
	return names, nil
}

func readUint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func readUint64(b []byte) uint64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
