// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package syscall

import (
	"sync"
	"unsafe"
)

type uintptr_t uint32
type size_t uint32

// Julien: do we actually want to expose all these types and syscalls ?
type Device_t uint64
type Fd_t uint32
type Fdflags_t uint32
type Filesize_t uint64
type Filetype_t uint8
type Inode_t uint64
type Linkcount_t uint64
type Lookupflags_t uint32
type Oflags_t uint32
type Rights_t uint64
type Timestamp_t uint64
type Dircookie_t uint64
type Filedelta_t int64
type Whence_t uint32
type Fstflags_t uint32

type Ciovec_t struct {
	buf     uintptr_t
	buf_len size_t
}

type Stat_t struct {
	Dev      Device_t
	Ino      Inode_t
	Filetype Filetype_t
	Nlink    Linkcount_t
	Size     Filesize_t
	Atime    Timestamp_t
	Mtime    Timestamp_t
	Ctime    Timestamp_t

	Mode int // FIXME

	// Uid and Gid are always zero on wasi platforms
	Uid uint32
	Gid uint32
}

type Fdstat_t struct {
	Filetype         Filetype_t
	Flags            Fdflags_t
	RightsBase       Rights_t
	RightsInheriting Rights_t
}

const (
	LOOKUP_SYMLINK_FOLLOW Lookupflags_t = 0x00000001

	OFLAG_CREATE    Oflags_t = 0x0001
	OFLAG_DIRECTORY Oflags_t = 0x0002
	OFLAG_EXCL      Oflags_t = 0x0004
	OFLAG_TRUNC     Oflags_t = 0x0008

	FDFLAG_APPEND   Fdflags_t = 0x0001
	FDFLAG_DSYNC    Fdflags_t = 0x0002
	FDFLAG_NONBLOCK Fdflags_t = 0x0004
	FDFLAG_RSYNC    Fdflags_t = 0x0008
	FDFLAG_SYNC     Fdflags_t = 0x0010

	RIGHT_FD_DATASYNC Rights_t = 1 << iota
	RIGHT_FD_READ
	RIGHT_FD_SEEK
	RIGHT_FDSTAT_SET_FLAGS
	RIGHT_FD_SYNC
	RIGHT_FD_TELL
	RIGHT_FD_WRITE
	RIGHT_FD_ADVISE
	RIGHT_FD_ALLOCATE
	RIGHT_PATH_CREATE_DIRECTORY
	RIGHT_PATH_CREATE_FILE
	RIGHT_PATH_LINK_SOURCE
	RIGHT_PATH_LINK_TARGET
	RIGHT_PATH_OPEN
	RIGHT_FD_READDIR
	RIGHT_PATH_READLINK
	RIGHT_PATH_RENAME_SOURCE
	RIGHT_PATH_RENAME_TARGET
	RIGHT_PATH_FILESTAT_GET
	RIGHT_PATH_FILESTAT_SET_SIZE
	RIGHT_PATH_FILESTAT_SET_TIMES
	RIGHT_FD_FILESTAT_GET
	RIGHT_FD_FILESTAT_SET_SIZE
	RIGHT_FD_FILESTAT_SET_TIMES
	RIGHT_PATH_SYMLINK
	RIGHT_PATH_REMOVE_DIRECTORY
	RIGHT_PATH_UNLINK_FILE
	RIGHT_POLL_FD_READWRITE
	RIGHT_SOCK_SHUTDOWN
	RIGHT_SOCK_ACCEPT

	RIGHT_FULL Rights_t = Rights_t(^uint32(0))

	// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-filetype-variant
	FILETYPE_UNKNOWN          Filetype_t = 0
	FILETYPE_BLOCK_DEVICE     Filetype_t = 1
	FILETYPE_CHARACTER_DEVICE Filetype_t = 2
	FILETYPE_DIRECTORY        Filetype_t = 3
	FILETYPE_REGULAR_FILE     Filetype_t = 4
	FILETYPE_SOCKET_DGRAM     Filetype_t = 5
	FILETYPE_SOCKET_STREAM    Filetype_t = 6
	FILETYPE_SYMBOLIC_LINK    Filetype_t = 7

	WHENCE_SET Whence_t = 0
	WHENCE_CUR Whence_t = 1
	WHENCE_END Whence_t = 2

	FILESTAT_SET_ATIM     Fstflags_t = 0x0001
	FILESTAT_SET_ATIM_NOW Fstflags_t = 0x0002
	FILESTAT_SET_MTIM     Fstflags_t = 0x0004
	FILESTAT_SET_MTIM_NOW Fstflags_t = 0x0008
)

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_closefd-fd---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_close
func Fd_close(
	fd Fd_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_filestat_set_sizefd-fd-size-filesize---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_filestat_set_size
func Fd_filestat_set_size(
	fd Fd_t,
	st_size Filesize_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_preadfd-fd-iovs-iovec_array-offset-filesize---resultsize-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_pread
func Fd_pread(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	offset Filesize_t,
	nread *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_pwrite
func Fd_pwrite(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	offset Filesize_t,
	nwritten *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_read
func Fd_read(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	nread *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_readdir
func Fd_readdir(
	fd Fd_t,
	buf *byte,
	buf_len size_t,
	cookie Dircookie_t,
	bufused *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_seek
func Fd_seek(
	fd Fd_t,
	offset Filedelta_t,
	whence Whence_t,
	newoffset *Filesize_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_get
func Fd_fdstat_get(
	fd Fd_t,
	buf *Fdstat_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_fdstat_set_rightsfd-fd-fs_rights_base-rights-fs_rights_inheriting-rights---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_rights
func Fd_fdstat_set_rights(
	fd Fd_t,
	rightsBase Rights_t,
	rightsInheriting Rights_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
func Fd_filestat_get(
	fd Fd_t,
	buf *Stat_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_write
func Fd_write(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	nwritten *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_create_directory
func Path_create_directory(
	fd Fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_get
func Path_filestat_get(
	fd Fd_t,
	flags Lookupflags_t,
	path *byte,
	path_len size_t,
	buf *Stat_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_set_times
func Path_filestat_set_times(
	fd Fd_t,
	flags Lookupflags_t,
	path *byte,
	path_len size_t,
	st_atim Timestamp_t,
	st_mtim Timestamp_t,
	fstflags Fstflags_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_link
func Path_link(
	old_fd Fd_t,
	old_flags Lookupflags_t,
	old_path *byte,
	old_path_len size_t,
	new_fd Fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_readlink
func Path_readlink(
	fd Fd_t,
	path *byte,
	path_len size_t,
	buf *byte,
	buf_len size_t,
	bufused *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_remove_directory
func Path_remove_directory(
	fd Fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_rename
func Path_rename(
	old_fd Fd_t,
	old_path *byte,
	old_path_len size_t,
	new_fd Fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_symlink
func Path_symlink(
	old_path *byte,
	old_path_len size_t,
	fd Fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_unlink_file
func Path_unlink_file(
	fd Fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_open
func Path_open(
	rootFD Fd_t,
	dirflags Lookupflags_t,
	path *byte,
	path_len size_t,
	oflags Oflags_t,
	fs_rights_base Rights_t,
	fs_rights_inheriting Rights_t,
	fs_flags Fdflags_t,
	fd *Fd_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 random_get
func Random_get(
	buf *byte,
	buf_len size_t,
) Errno

const rootFD Fd_t = 3 // TODO(Pryz): document where this magic number is coming from
var rootRightsDir Rights_t
var rootRightsFile Rights_t

var wd string

var fdPathsMu sync.Mutex
var fdPaths = make(map[int]string)

func init() {
	var stat Fdstat_t
	errno := Fd_fdstat_get(rootFD, &stat)
	if errno != 0 {
		// TODO(Pryz): if errno is EBADF it is likely because nothing
		// was mount into the module.
		panic("could not get fdstat of root: " + errno.Error())
	}
	rootRightsDir = stat.RightsBase
	rootRightsFile = stat.RightsInheriting

	wd, _ = Getenv("PWD")
}

// Provided by package runtime.
func now() (sec int64, nsec int32)

func isAbs(path string) bool {
	return len(path) > 0 && path[0] == '/'
}

func skipPathSeparator(path string) string {
	i := 0
	for i < len(path) && path[i] == '/' {
		i++
	}
	return path[i:]
}

func walkPath(path string) (root, tail string) {
	if isAbs(path) {
		return "/", skipPathSeparator(path)
	}
	for i < len(path) && path[i] != '/' {
		i++
	}
	return path[:i], skipPathSeparator(path[i:])
}

func dirPath(path string) string {
	i := len(path)
	for i > 0 && path[i-1] != '/' {
		i--
	}
	if i == 0 {
		return "/"
	}
	return path[:i]
}

func preparePath(path string, followTrailingSymlink bool) (*byte, size_t) {
	if path == "" || path[0] != '/' {
		path = wd + "/" + path
	}

	var resolvedPath string
	for {
		var part string
		part, path = walkPath(path)
		resolvedPath += "/" + part
		if i == len(parts)-1 && !followTrailingSymlink {
			break
		}
		for {
			dest, err := readlink("." + resolvedPath)
			if err != nil {
				break
			}
			if dest[0] != '/' {
				dest = dirPath(resolvedPath) + "/" + dest
			}
			resolvedPath = dest
		}
	}

	return &[]byte("." + resolvedPath)[0], size_t(1 + len(resolvedPath))
}

func readlink(path string) (string, error) {
	for buflen := size_t(128); ; buflen *= 2 {
		buf := make([]byte, buflen)
		var bufused size_t
		errno := Path_readlink(
			rootFD,
			&[]byte(path)[0],
			size_t(len(path)),
			&buf[0],
			buflen,
			&bufused,
		)
		if errno != 0 {
			return "", errnoErr(errno)
		}
		if bufused < buflen {
			return string(buf[:bufused]), nil
		}
	}
}

func Open(path string, openmode int, perm uint32) (int, error) {
	if path == "" {
		return 0, EINVAL
	}
	if path[0] != '/' {
		path = wd + "/" + path
	}

	path_ptr, path_len := preparePath(path, true)

	var oflags Oflags_t
	if openmode&O_CREATE != 0 {
		oflags |= OFLAG_CREATE
	}
	if openmode&O_TRUNC != 0 {
		oflags |= OFLAG_TRUNC
	}
	if openmode&O_EXCL != 0 {
		oflags |= OFLAG_EXCL
	}

	// Remove when https://github.com/bytecodealliance/wasmtime/pull/4967 is merged.
	st := &Stat_t{}
	if err := Stat(path, st); err != nil && err != ENOENT {
		return 0, err
	}
	if st.Filetype == FILETYPE_DIRECTORY {
		oflags |= OFLAG_DIRECTORY
	}

	var rights = rootRightsFile
	switch {
	case openmode&O_WRONLY != 0:
		rights &^= RIGHT_FD_READ | RIGHT_FD_READDIR
	case openmode&O_RDWR != 0:
		// no rights to remove
	default:
		rights &^= RIGHT_FD_DATASYNC | RIGHT_FD_WRITE | RIGHT_FD_ALLOCATE | RIGHT_PATH_FILESTAT_SET_SIZE
	}

	var fdflags Fdflags_t
	if openmode&O_APPEND != 0 {
		//fdflags |= FDFLAG_APPEN
	}
	if openmode&O_SYNC != 0 {
		fdflags |= FDFLAG_SYNC
	}

	rights = RIGHT_FULL

	var fd Fd_t
	errno := Path_open(
		rootFD,
		0,
		path_ptr,
		path_len,
		oflags,
		rights,
		rootRightsFile,
		fdflags,
		&fd,
	)

	fdPathsMu.Lock()
	fdPaths[int(fd)] = path
	fdPathsMu.Unlock()

	return int(fd), errnoErr(errno)
}

func Close(fd int) error {
	fdPathsMu.Lock()
	delete(fdPaths, fd)
	fdPathsMu.Unlock()

	errno := Fd_close(Fd_t(fd))
	return errnoErr(errno)
}

func CloseOnExec(fd int) {
	// nothing to do - no exec
}

func Mkdir(path string, perm uint32) error {
	path_ptr, path_len := preparePath(path, false)
	if errno := Path_create_directory(rootFD, path_ptr, path_len); errno != 0 {
		return errnoErr(errno)
	}
	// FIXME: matches rights to perm
	// Not all WASM runtime support rights so we ignore the potential error.
	_ = Fd_fdstat_set_rights(rootFD, RIGHT_FULL, RIGHT_FULL)
	return nil
}

func ReadDirent(fd int, buf []byte) (int, error) {
	return 0, ENOSYS
}

func ReadDir(fd int, buf []byte, cookie Dircookie_t) (int, error) {
	var bufused size_t
	errno := Fd_readdir(Fd_t(fd), &buf[0], size_t(len(buf)), cookie, &bufused)
	return int(bufused), errnoErr(errno)
}

func Stat(path string, st *Stat_t) error {
	path_ptr, path_len := preparePath(path, true)
	errno := Path_filestat_get(rootFD, 0, path_ptr, path_len, st)
	return errnoErr(errno)
}

func Lstat(path string, st *Stat_t) error {
	path_ptr, path_len := preparePath(path, false)
	errno := Path_filestat_get(rootFD, 0, path_ptr, path_len, st)
	return errnoErr(errno)
}

func Fstat(fd int, st *Stat_t) error {
	errno := Fd_filestat_get(Fd_t(fd), st)
	return errnoErr(errno)
}

func Unlink(path string) error {
	path_ptr, path_len := preparePath(path, false)
	errno := Path_unlink_file(rootFD, path_ptr, path_len)
	return errnoErr(errno)
}

func Rmdir(path string) error {
	path_ptr, path_len := preparePath(path, false)
	errno := Path_remove_directory(rootFD, path_ptr, path_len)
	return errnoErr(errno)
}

func Chmod(path string, mode uint32) error {
	return ENOSYS
}

func Fchmod(fd int, mode uint32) error {
	return ENOSYS
}

func Chown(path string, uid, gid int) error {
	return ENOSYS
}

func Fchown(fd int, uid, gid int) error {
	return ENOSYS
}

func Lchown(path string, uid, gid int) error {
	return ENOSYS
}

func UtimesNano(path string, ts []Timespec) error {
	path_ptr, path_len := preparePath(path, false)
	errno := Path_filestat_set_times(
		rootFD,
		0,
		path_ptr,
		path_len,
		Timestamp_t(TimespecToNsec(ts[0])),
		Timestamp_t(TimespecToNsec(ts[1])),
		FILESTAT_SET_ATIM|FILESTAT_SET_MTIM,
	)
	return errnoErr(errno)
}

func Rename(from, to string) error {
	old_path, old_path_len := preparePath(from, false)
	new_path, new_path_len := preparePath(to, false)
	errno := Path_rename(
		rootFD,
		old_path,
		old_path_len,
		rootFD,
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Truncate(path string, length int64) error {
	fd, openErr := Open(path, O_WRONLY, 0)
	if openErr != nil {
		return openErr
	}
	truncateErr := Ftruncate(fd, length)
	closeErr := Close(fd)
	if truncateErr != nil {
		return truncateErr
	}
	return closeErr
}

func Ftruncate(fd int, length int64) error {
	errno := Fd_filestat_set_size(Fd_t(fd), Filesize_t(length))
	return errnoErr(errno)
}

const ImplementsGetwd = true

func Getwd() (string, error) {
	return wd, nil
}

func Chdir(path string) (err error) {
	if path[0] != '/' {
		path = wd + "/" + path
	}
	var st Stat_t
	if err := Stat(path, &st); err != nil {
		return err
	}
	wd = path
	return nil
}

func Fchdir(fd int) error {
	fdPathsMu.Lock()
	wd = fdPaths[fd]
	fdPathsMu.Unlock()
	return nil
}

func Readlink(path string, buf []byte) (n int, err error) {
	path_ptr, path_len := preparePath(path, false)
	var bufused size_t
	errno := Path_readlink(
		rootFD,
		path_ptr,
		path_len,
		&buf[0],
		size_t(len(buf)),
		&bufused,
	)
	return int(bufused), errnoErr(errno)
}

func Link(path, link string) error {
	old_path, old_path_len := preparePath(path, false)
	new_path, new_path_len := preparePath(link, false)
	errno := Path_link(
		rootFD,
		0,
		old_path,
		old_path_len,
		rootFD,
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Symlink(path, link string) error {
	new_path, new_path_len := preparePath(link, false)
	errno := Path_symlink(
		&[]byte(path)[0],
		size_t(len(path)),
		rootFD, // TODO(Pryz): Meaning only works with absolute paths ?
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Fsync(fd int) error {
	return ENOSYS
}

func makeIOVec(b []byte) *Ciovec_t {
	return &Ciovec_t{
		buf:     uintptr_t(uintptr(unsafe.Pointer(&b[0]))),
		buf_len: size_t(len(b)),
	}
}

func Read(fd int, b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nread size_t
	errno := Fd_read(Fd_t(fd), makeIOVec(b), 1, &nread)
	return int(nread), errnoErr(errno)
}

func Write(fd int, b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nwritten size_t
	errno := Fd_write(Fd_t(fd), makeIOVec(b), 1, &nwritten)
	return int(nwritten), errnoErr(errno)
}

func Pread(fd int, b []byte, offset int64) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nread size_t
	errno := Fd_pread(Fd_t(fd), makeIOVec(b), 1, Filesize_t(offset), &nread)
	return int(nread), errnoErr(errno)
}

func Pwrite(fd int, b []byte, offset int64) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nwritten size_t
	errno := Fd_pwrite(Fd_t(fd), makeIOVec(b), 1, Filesize_t(offset), &nwritten)
	return int(nwritten), errnoErr(errno)
}

func Seek(fd int, offset int64, whence int) (int64, error) {
	var newoffset Filesize_t
	errno := Fd_seek(Fd_t(fd), Filedelta_t(offset), Whence_t(whence), &newoffset)
	return int64(newoffset), errnoErr(errno)
}

func Dup(fd int) (int, error) {
	return 0, ENOSYS
}

func Dup2(fd, newfd int) error {
	return ENOSYS
}

func Pipe(fd []int) error {
	return ENOSYS
}

func RandomGet(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	errno := Random_get(&b[0], size_t(len(b)))
	return errnoErr(errno)
}
