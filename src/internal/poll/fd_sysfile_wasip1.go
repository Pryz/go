// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package poll

import "syscall"

type SysFile struct {
	// Cache for the file type, lazily initialized when Seek is called.
	Filetype syscall.Filetype

	// If the file represents a directory, this field contains the current
	// readdir position. It is reset to zero if the program calls Seek(0, 0).
	Dircookie syscall.Dircookie

	// Absolute path of the file, as returned by syscall.PathOpen;
	// this is used by Fchdir to emulate setting the current directory
	// to an open file descriptor.
	Path string

	// TODO(achille): it could be meaningful to move isFile from FD to a method
	// on this struct type, and expose it as `IsFile() bool` which derives the
	// result from the Filetype field. We would need to ensure that Filetype is
	// always set instead of being lazily initialized.
}

// dupCloseOnExecOld always errors on wasip1 because there is no mechanism to
// duplicate file descriptors.
func dupCloseOnExecOld(fd int) (int, string, error) {
	return -1, "dup", syscall.ENOSYS
}
