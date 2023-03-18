// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package syscall

import (
	"sync"
	"unsafe"
)

type uintptr_t = uint32
type size_t = uint32
type __wasip1_fd_t = uint32
type __wasip1_fdflags_t = uint32
type __wasip1_filesize_t = uint64
type __wasip1_filetype_t = uint8
type __wasip1_lookupflags_t = uint32
type __wasip1_oflags_t = uint32
type __wasip1_rights_t = uint64
type __wasip1_timestamp_t = uint64
type __wasip1_dircookie_t = uint64
type __wasip1_filedelta_t = int64
type __wasip1_whence_t = uint32
type __wasip1_fstflags_t = uint32

type __wasip1_iovec_t struct {
	buf     uintptr_t
	buf_len size_t
}

type __wasip1_fdstat_t struct {
	Filetype         __wasip1_filetype_t
	Flags            __wasip1_fdflags_t
	RightsBase       __wasip1_rights_t
	RightsInheriting __wasip1_rights_t
}

const (
	LOOKUP_SYMLINK_FOLLOW = 0x00000001

	OFLAG_CREATE    = 0x0001
	OFLAG_DIRECTORY = 0x0002
	OFLAG_EXCL      = 0x0004
	OFLAG_TRUNC     = 0x0008

	FDFLAG_APPEND   = 0x0001
	FDFLAG_DSYNC    = 0x0002
	FDFLAG_NONBLOCK = 0x0004
	FDFLAG_RSYNC    = 0x0008
	FDFLAG_SYNC     = 0x0010

	RIGHT_FD_DATASYNC = 1 << iota
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

	RIGHT_FULL = ^uint64(0)

	WHENCE_SET = 0
	WHENCE_CUR = 1
	WHENCE_END = 2

	FILESTAT_SET_ATIM     = 0x0001
	FILESTAT_SET_ATIM_NOW = 0x0002
	FILESTAT_SET_MTIM     = 0x0004
	FILESTAT_SET_MTIM_NOW = 0x0008
)

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_closefd-fd---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_close
//go:noescape
func __wasip1_fd_close(fd __wasip1_fd_t) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_filestat_set_sizefd-fd-size-filesize---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_filestat_set_size
//go:noescape
func __wasip1_fd_filestat_set_size(
	fd __wasip1_fd_t,
	st_size __wasip1_filesize_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_preadfd-fd-iovs-iovec_array-offset-filesize---resultsize-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_pread
//go:noescape
func __wasip1_fd_pread(
	fd __wasip1_fd_t,
	iovs *__wasip1_iovec_t,
	iovs_len size_t,
	offset __wasip1_filesize_t,
	nread *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_pwrite
//go:noescape
func __wasip1_fd_pwrite(
	fd __wasip1_fd_t,
	iovs *__wasip1_iovec_t,
	iovs_len size_t,
	offset __wasip1_filesize_t,
	nwritten *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_read
//go:noescape
func __wasip1_fd_read(
	fd __wasip1_fd_t,
	iovs *__wasip1_iovec_t,
	iovs_len size_t,
	nread *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_readdir
//go:noescape
func __wasip1_fd_readdir(
	fd __wasip1_fd_t,
	buf *byte,
	buf_len size_t,
	cookie __wasip1_dircookie_t,
	bufused *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_seek
//go:noescape
func __wasip1_fd_seek(
	fd __wasip1_fd_t,
	offset __wasip1_filedelta_t,
	whence __wasip1_whence_t,
	newoffset *__wasip1_filesize_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_get
//go:noescape
func __wasip1_fd_fdstat_get(
	fd __wasip1_fd_t,
	buf *__wasip1_fdstat_t,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_fdstat_set_rightsfd-fd-fs_rights_base-rights-fs_rights_inheriting-rights---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_rights
//go:noescape
func __wasip1_fd_fdstat_set_rights(
	fd __wasip1_fd_t,
	rightsBase __wasip1_rights_t,
	rightsInheriting __wasip1_rights_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
//go:noescape
func __wasip1_fd_filestat_get(
	fd __wasip1_fd_t,
	buf *Stat_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func __wasip1_fd_write(
	fd __wasip1_fd_t,
	iovs *__wasip1_iovec_t,
	iovs_len size_t,
	nwritten *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_sync
//go:noescape
func __wasip1_fd_sync(fd __wasip1_fd_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_create_directory
//go:noescape
func __wasip1_path_create_directory(
	fd __wasip1_fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_get
//go:noescape
func __wasip1_path_filestat_get(
	fd __wasip1_fd_t,
	flags __wasip1_lookupflags_t,
	path *byte,
	path_len size_t,
	buf *Stat_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_set_times
//go:noescape
func __wasip1_path_filestat_set_times(
	fd __wasip1_fd_t,
	flags __wasip1_lookupflags_t,
	path *byte,
	path_len size_t,
	st_atim __wasip1_timestamp_t,
	st_mtim __wasip1_timestamp_t,
	fstflags __wasip1_fstflags_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_link
//go:noescape
func __wasip1_path_link(
	old_fd __wasip1_fd_t,
	old_flags __wasip1_lookupflags_t,
	old_path *byte,
	old_path_len size_t,
	new_fd __wasip1_fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_readlink
//go:noescape
func __wasip1_path_readlink(
	fd __wasip1_fd_t,
	path *byte,
	path_len size_t,
	buf *byte,
	buf_len size_t,
	bufused *size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_remove_directory
//go:noescape
func __wasip1_path_remove_directory(
	fd __wasip1_fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_rename
//go:noescape
func __wasip1_path_rename(
	old_fd __wasip1_fd_t,
	old_path *byte,
	old_path_len size_t,
	new_fd __wasip1_fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_symlink
//go:noescape
func __wasip1_path_symlink(
	old_path *byte,
	old_path_len size_t,
	fd __wasip1_fd_t,
	new_path *byte,
	new_path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_unlink_file
//go:noescape
func __wasip1_path_unlink_file(
	fd __wasip1_fd_t,
	path *byte,
	path_len size_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 path_open
//go:noescape
func __wasip1_path_open(
	rootFD __wasip1_fd_t,
	dirflags __wasip1_lookupflags_t,
	path *byte,
	path_len size_t,
	oflags __wasip1_oflags_t,
	fs_rights_base __wasip1_rights_t,
	fs_rights_inheriting __wasip1_rights_t,
	fs_flags __wasip1_fdflags_t,
	fd *__wasip1_fd_t,
) Errno

//go:wasmimport wasi_snapshot_preview1 random_get
//go:noescape
func __wasip1_random_get(
	buf *byte,
	buf_len size_t,
) Errno

type __wasip1_preopentype uint8

const (
	__wasip1_preopentype_dir __wasip1_preopentype = iota
)

type __wasip1_prestat_dir struct {
	pr_name_len size_t
}

type __wasip1_prestat struct {
	typ __wasip1_preopentype
	dir __wasip1_prestat_dir
}

//go:wasmimport wasi_snapshot_preview1 fd_prestat_get
//go:noescape
func __wasip1_fd_prestat_get(fd __wasip1_fd_t, prestat *__wasip1_prestat) Errno

//go:wasmimport wasi_snapshot_preview1 fd_prestat_dir_name
//go:noescape
func __wasip1_fd_prestat_dir_name(fd __wasip1_fd_t, path *byte, path_len size_t) Errno

var rootRightsDir __wasip1_rights_t
var rootRightsFile __wasip1_rights_t

var fdPathsMu sync.Mutex
var fdPaths = make(map[int]string)

type opendir struct {
	fd   __wasip1_fd_t
	name string
}

// List of preopen directories that were exposed by the runtime. The first one
// is assumed to the be root directory of the file system, and other others are
// seen as mount points at sub paths of the root.
var preopens []opendir

// File descriptor for the current working directory, defaults to the fd of the
// root directory which is expected to be the first preopen.
var cwd = opendir{
	fd: ^__wasip1_fd_t(0),
}

func init() {
	dirNameBuf := make([]byte, 256)
	// We start looking for preopens at fd=3 because 0, 1, and 2 are reserved
	// for standard input and outputs.
	for preopenFd := __wasip1_fd_t(3); ; preopenFd++ {
		var prestat __wasip1_prestat
		var errno = __wasip1_fd_prestat_get(preopenFd, &prestat)

		if errno == EBADF {
			break
		}
		if errno != 0 {
			panic("fd_prestat: " + errno.Error())
		}
		if prestat.typ != __wasip1_preopentype_dir {
			continue
		}
		if int(prestat.dir.pr_name_len) > len(dirNameBuf) {
			dirNameBuf = make([]byte, prestat.dir.pr_name_len)
		}

		errno = __wasip1_fd_prestat_dir_name(preopenFd, &dirNameBuf[0], prestat.dir.pr_name_len)
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
		var stat __wasip1_fdstat_t
		errno := __wasip1_fd_fdstat_get(preopens[0].fd, &stat)
		if errno != 0 {
			panic("fd_fdstat_get: " + errno.Error())
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
	i := 0
	for i < len(file) && file[i] == '/' {
		i++
	}
	file = file[i:]
	if dir == "/" {
		return dir + file
	}
	return dir + "/" + file
}

func preparePath(path string) (__wasip1_fd_t, string, *byte, size_t) {
	var dirfd __wasip1_fd_t
	var dirname string

	if len(path) == 0 || path[0] != '/' {
		dirfd = cwd.fd
		dirname = cwd.name
	} else if len(preopens) > 0 {
		dirfd = preopens[0].fd
		dirname = preopens[0].name
	} else {
		dirfd = ^__wasip1_fd_t(0)
		dirname = "/"
	}

	return dirfd, dirname, unsafe.StringData(path), size_t(len(path))
}

func Open(path string, openmode int, perm uint32) (int, error) {
	if path == "" {
		return 0, errEINVAL
	}

	var oflags __wasip1_oflags_t
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

	var fdflags __wasip1_fdflags_t
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

	var fd __wasip1_fd_t
	errno := __wasip1_path_open(
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

	// TODO(achille): this map is needed in order to support Fchdir and Getwd;
	// it's kind of bad, can we find an alternative?
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

	errno := __wasip1_fd_close(__wasip1_fd_t(fd))
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
	errno := __wasip1_path_create_directory(dirfd, path_ptr, path_len)
	return errnoErr(errno)
}

func ReadDir(fd int, buf []byte, cookie __wasip1_dircookie_t) (int, error) {
	var bufused size_t
	errno := __wasip1_fd_readdir(__wasip1_fd_t(fd), &buf[0], size_t(len(buf)), cookie, &bufused)
	return int(bufused), errnoErr(errno)
}

type Stat_t struct {
	Dev      uint64
	Ino      uint64
	Filetype uint8
	Nlink    uint64
	Size     uint64
	Atime    uint64
	Mtime    uint64
	Ctime    uint64

	Mode int

	// Uid and Gid are always zero on wasip1 platforms
	Uid uint32
	Gid uint32
}

func Stat(path string, st *Stat_t) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := __wasip1_path_filestat_get(dirfd, LOOKUP_SYMLINK_FOLLOW, path_ptr, path_len, st)
	setDefaultMode(st)
	return errnoErr(errno)
}

func Lstat(path string, st *Stat_t) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := __wasip1_path_filestat_get(dirfd, 0, path_ptr, path_len, st)
	setDefaultMode(st)
	return errnoErr(errno)
}

func Fstat(fd int, st *Stat_t) error {
	errno := __wasip1_fd_filestat_get(__wasip1_fd_t(fd), st)
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
	errno := __wasip1_path_unlink_file(dirfd, path_ptr, path_len)
	return errnoErr(errno)
}

func Rmdir(path string) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	errno := __wasip1_path_remove_directory(dirfd, path_ptr, path_len)
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
	errno := __wasip1_path_filestat_set_times(
		dirfd,
		LOOKUP_SYMLINK_FOLLOW,
		path_ptr,
		path_len,
		__wasip1_timestamp_t(TimespecToNsec(ts[0])),
		__wasip1_timestamp_t(TimespecToNsec(ts[1])),
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
	errno := __wasip1_path_rename(
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
	errno := __wasip1_fd_filestat_set_size(__wasip1_fd_t(fd), __wasip1_filesize_t(length))
	return errnoErr(errno)
}

const ImplementsGetwd = true

func Getwd() (string, error) {
	if cwd.fd == ^__wasip1_fd_t(0) {
		return "", errENOENT
	}
	return cwd.name, nil
}

func Chdir(path string) error {
	if path == "" {
		return errEINVAL
	}
	dirfd, dirname, path_ptr, path_len := preparePath(path)

	errno := __wasip1_path_open(
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
	__wasip1_fd_close(cwd.fd)
	cwd.fd = dirfd
	cwd.name = joinPath(dirname, path)
	return nil
}

func Fchdir(fd int) error {
	fdPathsMu.Lock()
	dir, ok := fdPaths[fd]
	fdPathsMu.Unlock()
	if !ok {
		return errEBADF
	}
	// wasi does not offer a mechanism to duplicate file descriptor so instead
	// emulate Fchdir by setting the current working directory to the path that
	// the file descriptor was originally opened from. This is necessary because
	// we don't have ownership of the file descriptor, which might be closed
	// after Fchdir returns and would break relataive path lookups if we had
	// retained it.
	return Chdir(dir)
}

func Readlink(path string, buf []byte) (n int, err error) {
	if path == "" {
		return 0, errEINVAL
	}
	dirfd, _, path_ptr, path_len := preparePath(path)
	var bufused size_t
	errno := __wasip1_path_readlink(
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
	errno := __wasip1_path_link(
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
	errno := __wasip1_path_symlink(
		&[]byte(path)[0],
		size_t(len(path)),
		dirfd,
		new_path,
		new_path_len,
	)
	return errnoErr(errno)
}

func Fsync(fd int) error {
	errno := __wasip1_fd_sync(__wasip1_fd_t(fd))
	return errnoErr(errno)
}

func makeIOVec(b []byte) *__wasip1_iovec_t {
	return &__wasip1_iovec_t{
		buf:     uintptr_t(uintptr(unsafe.Pointer(&b[0]))),
		buf_len: size_t(len(b)),
	}
}

func Read(fd int, b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nread size_t
	errno := __wasip1_fd_read(__wasip1_fd_t(fd), makeIOVec(b), 1, &nread)
	return int(nread), errnoErr(errno)
}

func Write(fd int, b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nwritten size_t
	errno := __wasip1_fd_write(__wasip1_fd_t(fd), makeIOVec(b), 1, &nwritten)
	return int(nwritten), errnoErr(errno)
}

func Pread(fd int, b []byte, offset int64) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nread size_t
	errno := __wasip1_fd_pread(__wasip1_fd_t(fd), makeIOVec(b), 1, __wasip1_filesize_t(offset), &nread)
	return int(nread), errnoErr(errno)
}

func Pwrite(fd int, b []byte, offset int64) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var nwritten size_t
	errno := __wasip1_fd_pwrite(__wasip1_fd_t(fd), makeIOVec(b), 1, __wasip1_filesize_t(offset), &nwritten)
	return int(nwritten), errnoErr(errno)
}

func Seek(fd int, offset int64, whence int) (int64, error) {
	var newoffset __wasip1_filesize_t
	errno := __wasip1_fd_seek(__wasip1_fd_t(fd), __wasip1_filedelta_t(offset), __wasip1_whence_t(whence), &newoffset)
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
	errno := __wasip1_random_get(&b[0], size_t(len(b)))
	return errnoErr(errno)
}
