// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

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

	Mode int

	// Uid and Gid are always zero on wasip1 platforms
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
//go:noescape
func Fd_close(
	fd Fd_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_filestat_set_sizefd-fd-size-filesize---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_filestat_set_size
//go:noescape
func Fd_filestat_set_size(
	fd Fd_t,
	st_size Filesize_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_preadfd-fd-iovs-iovec_array-offset-filesize---resultsize-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_pread
//go:noescape
func Fd_pread(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	offset Filesize_t,
	nread *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_pwrite
//go:noescape
func Fd_pwrite(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	offset Filesize_t,
	nwritten *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_read
//go:noescape
func Fd_read(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	nread *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_readdir
//go:noescape
func Fd_readdir(
	fd Fd_t,
	buf *byte,
	buf_len size_t,
	cookie Dircookie_t,
	bufused *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_seek
//go:noescape
func Fd_seek(
	fd Fd_t,
	offset Filedelta_t,
	whence Whence_t,
	newoffset *Filesize_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_get
//go:noescape
func Fd_fdstat_get(
	fd Fd_t,
	buf *Fdstat_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_fdstat_set_rightsfd-fd-fs_rights_base-rights-fs_rights_inheriting-rights---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_rights
//go:noescape
func Fd_fdstat_set_rights(
	fd Fd_t,
	rightsBase Rights_t,
	rightsInheriting Rights_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
//go:noescape
func Fd_filestat_get(
	fd Fd_t,
	buf *Stat_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func Fd_write(
	fd Fd_t,
	iovs *Ciovec_t,
	iovs_len size_t,
	nwritten *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_sync
//go:noescape
func Fd_sync(fd Fd_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_create_directory
//go:noescape
func Path_create_directory(
	fd Fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_get
//go:noescape
func Path_filestat_get(
	fd Fd_t,
	flags Lookupflags_t,
	path *byte,
	path_len size_t,
	buf *Stat_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_set_times
//go:noescape
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
//go:noescape
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
//go:noescape
func Path_readlink(
	fd Fd_t,
	path *byte,
	path_len size_t,
	buf *byte,
	buf_len size_t,
	bufused *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_remove_directory
//go:noescape
func Path_remove_directory(
	fd Fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_rename
//go:noescape
func Path_rename(
	old_fd Fd_t,
	old_path *byte,
	old_path_len size_t,
	new_fd Fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_symlink
//go:noescape
func Path_symlink(
	old_path *byte,
	old_path_len size_t,
	fd Fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_unlink_file
//go:noescape
func Path_unlink_file(
	fd Fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_open
//go:noescape
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
//go:noescape
func Random_get(
	buf *byte,
	buf_len size_t,
) Errno

type __wasi_preopentype uint8

const (
	__wasi_preopentype_dir __wasi_preopentype = iota
)

type __wasi_prestat_dir struct {
	pr_name_len size_t
}

type __wasi_prestat struct {
	typ __wasi_preopentype
	dir __wasi_prestat_dir
}

//go:wasmimport wasi_snapshot_preview1 fd_prestat_get
//go:noescape
func __wasi_fd_prestat_get(fd Fd_t, prestat *__wasi_prestat) Errno

//go:wasmimport wasi_snapshot_preview1 fd_prestat_dir_name
//go:noescape
func __wasi_fd_prestat_dir_name(fd Fd_t, path *byte, path_len size_t) Errno

// const rootFD Fd_t = 3 // TODO(Pryz): document where this magic number is coming from
var rootRightsDir Rights_t
var rootRightsFile Rights_t

//var wd string

var fdPathsMu sync.Mutex
var fdPaths = make(map[int]string)

type opendir struct {
	fd   Fd_t
	name string
}

// List of preopen directories that were exposed by the runtime. The first one
// is assumed to the be root directory of the file system, and other others are
// seen as mount points at sub paths of the root.
var preopens []opendir

// File descriptor for the current working directory, defaults to the fd of the
// root directory which is expected to be the first preopen.
var cwd = opendir{
	fd: ^Fd_t(0),
}

func init() {
	dirNameBuf := make([]byte, 256)
	// We start looking for preopens at fd=3 because 0, 1, and 2 are reserved
	// for standard input and outputs.
	for preopenFd := Fd_t(3); ; preopenFd++ {
		var prestat __wasi_prestat
		var errno = __wasi_fd_prestat_get(preopenFd, &prestat)

		if errno == EBADF {
			break
		}
		if errno != 0 {
			panic("fd_prestat: " + errno.Error())
		}
		if prestat.typ != __wasi_preopentype_dir {
			continue
		}
		if int(prestat.dir.pr_name_len) > len(dirNameBuf) {
			dirNameBuf = make([]byte, prestat.dir.pr_name_len)
		}

		errno = __wasi_fd_prestat_dir_name(preopenFd, &dirNameBuf[0], prestat.dir.pr_name_len)
		if errno != 0 {
			panic("fd_prestat_dir_name: " + errno.Error())
			continue
		}

		preopens = append(preopens, opendir{
			fd:   preopenFd,
			name: string(dirNameBuf[:prestat.dir.pr_name_len]),
		})
	}

	if len(preopens) > 0 {
		var stat Fdstat_t
		errno := Fd_fdstat_get(preopens[0].fd, &stat)
		if errno != 0 {
			// TODO(Pryz): if errno is EBADF it is likely because nothing
			// was mount into the module.
			panic("could not get fdstat of root: " + errno.Error())
		}
		cwd = preopens[0]
		rootRightsDir = stat.RightsBase
		rootRightsFile = stat.RightsInheriting
	}

	pwd, _ := Getenv("PWD")
	Chdir(pwd)
}

// Provided by package runtime.
func now() (sec int64, nsec int32)

func joinPath(dir, file string) string {
	if dir == "/" {
		return dir + file
	}
	return dir + "/" + file
}

func preparePath(path string) (Fd_t, string, *byte, size_t) {
	var dirfd Fd_t
	var dirname string

	if len(path) == 0 || path[0] != '/' {
		dirfd = cwd.fd
		dirname = cwd.name
	} else if len(preopens) > 0 {
		dirfd = preopens[0].fd
		dirname = preopens[0].name
	} else {
		dirfd = ^Fd_t(0)
		dirname = "/"
	}

	return dirfd, dirname, unsafe.StringData(path), size_t(len(path))
}

func relativeDirFd(path string) Fd_t {
	// TODO(achille): we might want to look for the last preopen which is
	// a prefix of the path received as argument.
	if len(path) > 0 && path[0] == '/' && len(preopens) > 0 {
		return preopens[0].fd
	}
	return cwd.fd
}

func Open(path string, openmode int, perm uint32) (int, error) {
	if path == "" {
		return 0, errEINVAL
	}

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
	// st := &Stat_t{}
	// if err := Stat(path, st); err != nil && err != ENOENT {
	// 	return 0, err
	// }
	// if st.Filetype == FILETYPE_DIRECTORY {
	// 	oflags |= OFLAG_DIRECTORY
	// }

	var rights = rootRightsFile
	switch {
	case openmode&O_WRONLY != 0:
		rights &^= RIGHT_FD_READ | RIGHT_FD_READDIR
		// TODO(achille): wazero needs to offer a mechanism for setting write
		// permissions on open files; at this time there is none so we add the
		// O_CREATE flag to force it.
		oflags |= OFLAG_CREATE
	case openmode&O_RDWR != 0:
		// no rights to remove
		oflags |= OFLAG_CREATE
	default:
		rights &^= RIGHT_FD_DATASYNC | RIGHT_FD_WRITE | RIGHT_FD_ALLOCATE | RIGHT_PATH_FILESTAT_SET_SIZE
	}

	var fdflags Fdflags_t
	if openmode&O_APPEND != 0 {
		// TODO(achille): why was this commented out?
		fdflags |= FDFLAG_APPEND
	}
	if openmode&O_SYNC != 0 {
		fdflags |= FDFLAG_SYNC
	}

	// TODO(achille): decide if we set rights, and if we don't we should
	// remove the code above.
	rights = RIGHT_FULL

	dirfd, dirname, path_ptr, path_len := preparePath(path)

	var fd Fd_t
	errno := Path_open(
		dirfd,
		LOOKUP_SYMLINK_FOLLOW,
		path_ptr,
		path_len,
		oflags,
		rights,
		rootRightsFile,
		fdflags,
		&fd,
	)

	// TODO(achille): this map is needed in order to support Fchdir and
	// Getwd; it's kind of bad, can we find an alternative?
	path = joinPath(dirname, path)
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
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	if errno := Path_create_directory(dirfd, path_ptr, path_len); errno != 0 {
		return errnoErr(errno)
	}
	// TODO(achille): I commented this code because I don't think that passing
	// rootFD is the right thing to do here.
	//
	// FIXME: matches rights to perm
	// Not all WASM runtime support rights so we ignore the potential error.
	// _ = Fd_fdstat_set_rights(rootFD, RIGHT_FULL, RIGHT_FULL)
	return nil
}

func ReadDir(fd int, buf []byte, cookie Dircookie_t) (int, error) {
	var bufused size_t
	errno := Fd_readdir(Fd_t(fd), &buf[0], size_t(len(buf)), cookie, &bufused)
	return int(bufused), errnoErr(errno)
}

func Stat(path string, st *Stat_t) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := Path_filestat_get(dirfd, LOOKUP_SYMLINK_FOLLOW, path_ptr, path_len, st)
	setDefaultMode(st)
	return errnoErr(errno)
}

func Lstat(path string, st *Stat_t) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := Path_filestat_get(dirfd, 0, path_ptr, path_len, st)
	setDefaultMode(st)
	return errnoErr(errno)
}

func Fstat(fd int, st *Stat_t) error {
	errno := Fd_filestat_get(Fd_t(fd), st)
	setDefaultMode(st)
	return errnoErr(errno)
}

func setDefaultMode(st *Stat_t) {
	// WASI does not support unix-like permissions, but Go programs are likely
	// to expect the permission bits to not be zero so we set defaults to help
	// avoid breaking applications that are migrating to WASM.
	if st.Filetype == FILETYPE_DIRECTORY {
		st.Mode = 0700
	} else {
		st.Mode = 0600
	}
}

func Unlink(path string) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := Path_unlink_file(dirfd, path_ptr, path_len)
	return errnoErr(errno)
}

func Rmdir(path string) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := Path_remove_directory(dirfd, path_ptr, path_len)
	return errnoErr(errno)
}

func Chmod(path string, mode uint32) error {
	var stat Stat_t
	return Stat(path, &stat)
}

func Fchmod(fd int, mode uint32) error {
	var stat Stat_t
	return Fstat(fd, &stat)
}

func Chown(path string, uid, gid int) error {
	return errENOSYS
}

func Fchown(fd int, uid, gid int) error {
	return errENOSYS
}

func Lchown(path string, uid, gid int) error {
	return errENOSYS
}

func UtimesNano(path string, ts []Timespec) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := Path_filestat_set_times(
		dirfd,
		LOOKUP_SYMLINK_FOLLOW,
		path_ptr,
		path_len,
		Timestamp_t(TimespecToNsec(ts[0])),
		Timestamp_t(TimespecToNsec(ts[1])),
		FILESTAT_SET_ATIM|FILESTAT_SET_MTIM,
	)
	return errnoErr(errno)
}

func Rename(from, to string) error {
	if from == "" || to == "" {
		return errEINVAL
	}
	old_dirfd, _, old_path, old_path_len := preparePath(from)
	new_dirfd, _, new_path, new_path_len := preparePath(to)
	errno := Path_rename(
		old_dirfd,
		old_path,
		old_path_len,
		new_dirfd,
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Truncate(path string, length int64) error {
	if path == "" {
		return errEINVAL
	}
	// We use O_APPEND here because it is the only way to get wazero to set the
	// O_RDWR open flag on the open file, which is needed to truncate ta file,
	// see ftruncate(2):
	//
	//  [EINVAL] fildes is not open for writing.
	//
	fd, err := Open(path, O_WRONLY|O_APPEND, 0)
	if err != nil {
		return err
	}
	defer Close(fd)
	return Ftruncate(fd, length)
}

func Ftruncate(fd int, length int64) error {
	errno := Fd_filestat_set_size(Fd_t(fd), Filesize_t(length))
	return errnoErr(errno)
}

const ImplementsGetwd = true

func Getwd() (string, error) {
	if cwd.fd == ^Fd_t(0) {
		return "", errENOENT
	}
	return cwd.name, nil
}

func Chdir(path string) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, dirname, path_ptr, path_len := preparePath(path)

	errno := Path_open(
		dirfd,
		LOOKUP_SYMLINK_FOLLOW,
		path_ptr,
		path_len,
		OFLAG_DIRECTORY,
		RIGHT_FULL,
		RIGHT_FULL,
		0,
		&dirfd,
	)

	if err := errnoErr(errno); err != nil {
		return err
	}

	// Note: this might be a preopen, we rely on the runtime to forbid closing
	// of preopen file descriptors (this is validated by wasi-testsuite).
	Fd_close(cwd.fd)
	cwd.fd = dirfd
	cwd.name = joinPath(dirname, path)
	return nil
}

func Fchdir(fd int) error {
	fdPathsMu.Lock()
	defer fdPathsMu.Unlock()

	dir, ok := fdPaths[fd]
	if !ok {
		return errEBADF
	}

	cwd.fd = Fd_t(fd)
	cwd.name = dir
	return nil
}

func Readlink(path string, buf []byte) (n int, err error) {
	if path == "" {
		return 0, errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	var bufused size_t
	errno := Path_readlink(
		dirfd,
		path_ptr,
		path_len,
		&buf[0],
		size_t(len(buf)),
		&bufused,
	)
	return int(bufused), errnoErr(errno)
}

func Link(path, link string) error {
	if path == "" || link == "" {
		return errEINVAL
	}
	old_dirfd, _, old_path, old_path_len := preparePath(path)
	new_dirfd, _, new_path, new_path_len := preparePath(link)
	errno := Path_link(
		old_dirfd,
		LOOKUP_SYMLINK_FOLLOW,
		old_path,
		old_path_len,
		new_dirfd,
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Symlink(path, link string) error {
	if path == "" || link == "" {
		return errEINVAL
	}
	dirfd, _, new_path, new_path_len := preparePath(link)
	errno := Path_symlink(
		&[]byte(path)[0],
		size_t(len(path)),
		dirfd,
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Fsync(fd int) error {
	errno := Fd_sync(Fd_t(fd))
	return errnoErr(errno)
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
	return 0, errENOSYS
}

func Dup2(fd, newfd int) error {
	return errENOSYS
}

func Pipe(fd []int) error {
	return errENOSYS
}

func RandomGet(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	errno := Random_get(&b[0], size_t(len(b)))
	return errnoErr(errno)
}
