// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package poll

import (
	"syscall"
	"unsafe"
)

// ReadDir wraps syscall.ReadDir.
// We treat this like an ordinary system call rather than a call
// that tries to fill the buffer.
func (fd *FD) ReadDir(buf []byte, cookie syscall.Dircookie) (int, error) {
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

func (fd *FD) ReadDirent(buf []byte) (int, error) {
	n, err := fd.ReadDir(buf, fd.Dircookie)
	if err != nil {
		return 0, err
	}
	if n <= 0 {
		return n, nil // EOF
	}

	// We assume that the caller of ReadDirent will consume the entire buffer
	// up to the last full entry, so we scan through the buffer looking for the
	// value of the last next cookie.
	b := buf[:n]

	for len(b) > 0 {
		next, ok := direntNext(b)
		if !ok {
			break
		}
		size, ok := direntReclen(b)
		if !ok {
			break
		}
		if size > uint64(len(b)) {
			break
		}
		fd.Dircookie = syscall.Dircookie(next)
		b = b[size:]
	}

	// Trim a potentially incomplete trailing entry; this is necessary because
	// the code in src/os/dir_unix.go does not deal well with partial values in
	// calls to direntReclen, etc... and ends up causing an early EOF before all
	// directory entries were consumed. ReadDirent is called with a large enough
	// buffer (8 KiB) that at least one entry should always fit, tho this seems
	// a bit brittle but cannot be addressed without a large change of the
	// algorithm in the os.(*File).readdir method.
	return n - len(b), nil
}

func direntReclen(buf []byte) (uint64, bool) {
	namelen, ok := direntNamlen(buf)
	return 24 + namelen, ok
}

func direntNamlen(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(syscall.Dirent{}.Namlen), unsafe.Sizeof(syscall.Dirent{}.Namlen))
}

func direntNext(buf []byte) (uint64, bool) {
	return readInt(buf, unsafe.Offsetof(syscall.Dirent{}.Next), unsafe.Sizeof(syscall.Dirent{}.Next))
}

// readInt returns the size-bytes unsigned integer in native byte order at offset off.
func readInt(b []byte, off, size uintptr) (u uint64, ok bool) {
	if len(b) < int(off+size) {
		return 0, false
	}
	return readIntLE(b[off:], size), true
}

func readIntBE(b []byte, size uintptr) uint64 {
	switch size {
	case 1:
		return uint64(b[0])
	case 2:
		_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
		return uint64(b[1]) | uint64(b[0])<<8
	case 4:
		_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
		return uint64(b[3]) | uint64(b[2])<<8 | uint64(b[1])<<16 | uint64(b[0])<<24
	case 8:
		_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
		return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
			uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	default:
		panic("syscall: readInt with unsupported size")
	}
}

func readIntLE(b []byte, size uintptr) uint64 {
	switch size {
	case 1:
		return uint64(b[0])
	case 2:
		_ = b[1] // bounds check hint to compiler; see golang.org/issue/14808
		return uint64(b[0]) | uint64(b[1])<<8
	case 4:
		_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24
	case 8:
		_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
		return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
			uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	default:
		panic("syscall: readInt with unsupported size")
	}
}
