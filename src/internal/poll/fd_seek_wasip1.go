// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package poll

import (
	"sync/atomic"
	"syscall"
)

// Seek wraps syscall.Seek.
func (fd *FD) Seek(offset int64, whence int) (int64, error) {
	if err := fd.incref(); err != nil {
		return 0, err
	}
	defer fd.decref()
	fileType := atomic.LoadUint32(&fd.Filetype)

	if fileType == syscall.FILETYPE_UNKNOWN {
		var stat syscall.Stat_t
		if err := fd.Fstat(&stat); err != nil {
			return 0, err
		}
		fileType = syscall.Filetype(stat.Filetype)
		atomic.StoreUint32(&fd.Filetype, fileType)
	}

	if fileType == syscall.FILETYPE_DIRECTORY {
		// If the file descriptor is opened on a directory, we reset the readdir
		// cookie when seeking back to the beginning to allow reusing the file
		// descriptor to scan the directory again.
		if offset == 0 && whence == 0 {
			fd.Dircookie = 0
			return 0, nil
		} else {
			return 0, syscall.EINVAL
		}
	}

	return syscall.Seek(fd.Sysfd, offset, whence)
}
