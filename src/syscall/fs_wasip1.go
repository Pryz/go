// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package syscall

import (
	"runtime"
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
)

const (
	OFLAG_CREATE    = 0x0001
	OFLAG_DIRECTORY = 0x0002
	OFLAG_EXCL      = 0x0004
	OFLAG_TRUNC     = 0x0008
)

const (
	FDFLAG_APPEND   = 0x0001
	FDFLAG_DSYNC    = 0x0002
	FDFLAG_NONBLOCK = 0x0004
	FDFLAG_RSYNC    = 0x0008
	FDFLAG_SYNC     = 0x0010
)

const (
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
)

const (
	WHENCE_SET = 0
	WHENCE_CUR = 1
	WHENCE_END = 2
)

const (
	FILESTAT_SET_ATIM     = 0x0001
	FILESTAT_SET_ATIM_NOW = 0x0002
	FILESTAT_SET_MTIM     = 0x0004
	FILESTAT_SET_MTIM_NOW = 0x0008
)

const (
	// Despite the rights being defined as a 64 bits integer in the spec,
	// wasmtime crashes the program if we set any of the upper 32 bits.
	fullRights  = __wasip1_rights_t(^uint32(0))
	readRights  = __wasip1_rights_t(RIGHT_FD_READ | RIGHT_FD_READDIR)
	writeRights = __wasip1_rights_t(RIGHT_FD_DATASYNC | RIGHT_FD_WRITE | RIGHT_FD_ALLOCATE | RIGHT_PATH_FILESTAT_SET_SIZE)
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
func __wasip1_fd_filestat_set_size(fd __wasip1_fd_t, st_size __wasip1_filesize_t) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_preadfd-fd-iovs-iovec_array-offset-filesize---resultsize-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_pread
//go:noescape
func __wasip1_fd_pread(fd __wasip1_fd_t, iovs *__wasip1_iovec_t, iovs_len size_t, offset __wasip1_filesize_t, nread *size_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_pwrite
//go:noescape
func __wasip1_fd_pwrite(fd __wasip1_fd_t, iovs *__wasip1_iovec_t, iovs_len size_t, offset __wasip1_filesize_t, nwritten *size_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_read
//go:noescape
func __wasip1_fd_read(fd __wasip1_fd_t, iovs *__wasip1_iovec_t, iovs_len size_t, nread *size_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_readdir
//go:noescape
func __wasip1_fd_readdir(fd __wasip1_fd_t, buf *byte, buf_len size_t, cookie __wasip1_dircookie_t, bufused *size_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_seek
//go:noescape
func __wasip1_fd_seek(fd __wasip1_fd_t, offset __wasip1_filedelta_t, whence __wasip1_whence_t, newoffset *__wasip1_filesize_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_get
//go:noescape
func __wasip1_fd_fdstat_get(fd __wasip1_fd_t, buf *__wasip1_fdstat_t) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_fdstat_set_rightsfd-fd-fs_rights_base-rights-fs_rights_inheriting-rights---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_rights
//go:noescape
func __wasip1_fd_fdstat_set_rights(fd __wasip1_fd_t, rightsBase __wasip1_rights_t, rightsInheriting __wasip1_rights_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
//go:noescape
func __wasip1_fd_filestat_get(fd __wasip1_fd_t, buf *Stat_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func __wasip1_fd_write(fd __wasip1_fd_t, iovs *__wasip1_iovec_t, iovs_len size_t, nwritten *size_t) Errno

//go:wasmimport wasi_snapshot_preview1 fd_sync
//go:noescape
func __wasip1_fd_sync(fd __wasip1_fd_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_create_directory
//go:noescape
func __wasip1_path_create_directory(fd __wasip1_fd_t, path *byte, path_len size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_get
//go:noescape
func __wasip1_path_filestat_get(fd __wasip1_fd_t, flags __wasip1_lookupflags_t, path *byte, path_len size_t, buf *Stat_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_filestat_set_times
//go:noescape
func __wasip1_path_filestat_set_times(fd __wasip1_fd_t, flags __wasip1_lookupflags_t, path *byte, path_len size_t, st_atim __wasip1_timestamp_t, st_mtim __wasip1_timestamp_t, fstflags __wasip1_fstflags_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_link
//go:noescape
func __wasip1_path_link(old_fd __wasip1_fd_t, old_flags __wasip1_lookupflags_t, old_path *byte, old_path_len size_t, new_fd __wasip1_fd_t, new_path *byte, new_path_len size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_readlink
//go:noescape
func __wasip1_path_readlink(fd __wasip1_fd_t, path *byte, path_len size_t, buf *byte, buf_len size_t, bufused *size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_remove_directory
//go:noescape
func __wasip1_path_remove_directory(fd __wasip1_fd_t, path *byte, path_len size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_rename
//go:noescape
func __wasip1_path_rename(old_fd __wasip1_fd_t, old_path *byte, old_path_len size_t, new_fd __wasip1_fd_t, new_path *byte, new_path_len size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_symlink
//go:noescape
func __wasip1_path_symlink(old_path *byte, old_path_len size_t, fd __wasip1_fd_t, new_path *byte, new_path_len size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_unlink_file
//go:noescape
func __wasip1_path_unlink_file(fd __wasip1_fd_t, path *byte, path_len size_t) Errno

//go:wasmimport wasi_snapshot_preview1 path_open
//go:noescape
func __wasip1_path_open(rootFD __wasip1_fd_t, dirflags __wasip1_lookupflags_t, path *byte, path_len size_t, oflags __wasip1_oflags_t, fs_rights_base __wasip1_rights_t, fs_rights_inheriting __wasip1_rights_t, fs_flags __wasip1_fdflags_t, fd *__wasip1_fd_t) Errno

//go:wasmimport wasi_snapshot_preview1 random_get
//go:noescape
func __wasip1_random_get(buf *byte, buf_len size_t) Errno

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

type opendir struct {
	fd   __wasip1_fd_t
	name string
}

// List of preopen directories that were exposed by the runtime. The first one
// is assumed to the be root directory of the file system, and others are seen
// as mount points at sub paths of the root.
var preopens []opendir

// Current working directory. We maintain this as a string and resolve paths in
// the code because wasmtime does not allow relative path lookups outside of the
// scope of a directory; a previous approach we tried consisted in maintaining
// open a file descriptor to the current directory so we could perform relative
// path lookups from that location, but it resulted in breaking path resolution
// from the current directory to its parent.
var cwd string

func init() {
	dirNameBuf := make([]byte, 256)
	// We start looking for preopens at fd=3 because 0, 1, and 2 are reserved
	// for standard input and outputs.
	for preopenFd := __wasip1_fd_t(3); ; preopenFd++ {
		var prestat __wasip1_prestat

		errno := __wasip1_fd_prestat_get(preopenFd, &prestat)
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

	if cwd, _ = Getenv("PWD"); cwd != "" {
		cwd = joinPath("/", cwd)
	} else if len(preopens) > 0 {
		cwd = preopens[0].name
	}
}

// Provided by package runtime.
func now() (sec int64, nsec int32)

//go:nosplit
func appendCleanPath(buf []byte, path string, lookupParent bool) ([]byte, bool) {
	i := 0
	for i < len(path) {
		for i < len(path) && path[i] == '/' {
			i++
		}

		j := i
		for j < len(path) && path[j] != '/' {
			j++
		}

		s := path[i:j]
		i = j

		switch s {
		case "":
			continue
		case ".":
			continue
		case "..":
			if !lookupParent {
				k := len(buf)
				for k > 0 && buf[k-1] != '/' {
					k--
				}
				for k > 1 && buf[k-1] == '/' {
					k--
				}
				buf = buf[:k]
				if k == 0 {
					lookupParent = true
				} else {
					s = ""
					continue
				}
			}
		default:
			lookupParent = false
		}

		if len(buf) > 0 && buf[len(buf)-1] != '/' {
			buf = append(buf, '/')
		}
		buf = append(buf, s...)
	}
	return buf, lookupParent
}

// joinPath concatenates dir and file paths, producing a cleaned path where
// "." and ".." have been removed, unless dir is relative and the references
// to parent directories in file represented a location relatie to a parent
// of dir.
//
// This function is used for path resolution of all wasi functions expecting
// a path argument; the returned string is heap allocated, which we may want
// to optimize in the future. Instead of returning a string, the function
// could append the result to an output buffer that the functions in this
// file can manage to have allocated on the stack (e.g. initializing to a
// fixed capacity). Since it will significantly increase code complexity,
// we prefer to optimize for readability and maintainability at this time.
func joinPath(dir, file string) string {
	buf := make([]byte, 0, len(dir)+len(file)+1)
	if isAbs(dir) {
		buf = append(buf, '/')
	}
	buf, lookupParent := appendCleanPath(buf, dir, false)
	buf, _ = appendCleanPath(buf, file, lookupParent)
	// The appendCleanPath function cleans the path so it does not inject
	// references to the current directory. If both the dir and file args
	// were ".", this results in the output buffer being empty so we handle
	// this condition here.
	if len(buf) == 0 {
		buf = append(buf, '.')
	}
	// If the file ended with a '/' we make sure that the output also ends
	// with a '/'. This is needed to ensure that programs have a mechanism
	// to represent dereferencing symbolic links pointing to directories.
	if buf[len(buf)-1] != '/' && isDir(file) {
		buf = append(buf, '/')
	}
	return unsafe.String(&buf[0], len(buf))
}

func isAbs(path string) bool {
	return hasPrefix(path, "/")
}

func isDir(path string) bool {
	return hasSuffix(path, "/")
}

func hasPrefix(s, p string) bool {
	return len(s) >= len(p) && s[:len(p)] == p
}

func hasSuffix(s, x string) bool {
	return len(s) >= len(x) && s[len(s)-len(x):] == x
}

// preparePath returns the preopen file descriptor and name of the directory to
// perform path resolution from, along with the pair of pointer and length for
// the relative expression of path from the directory.
//
// If the path argument is not absolute, it is first appended to the current
// working directory before resolution.
func preparePath(path string) (__wasip1_fd_t, string, *byte, size_t) {
	var dirFd = ^__wasip1_fd_t(0)
	var dirName string

	dir := "/"
	if !isAbs(path) {
		dir = cwd
	}
	path = joinPath(dir, path)

	for _, p := range preopens {
		if len(p.name) > len(dirName) && hasPrefix(path, p.name) {
			dirFd, dirName = p.fd, p.name
		}
	}

	path = path[len(dirName):]
	for isAbs(path) {
		path = path[1:]
	}
	if len(path) == 0 {
		path = "."
	}

	return dirFd, dirName, unsafe.StringData(path), size_t(len(path))
}

func Open(path string, openmode int, perm uint32) (int, error) {
	fd, _, err := PathOpen(path, openmode)
	return fd, err
}

// PathOpen is similar to Open but it also returns the absolute path of the
// newly open file, which may differ from the path passed as argument if it was
// relative and therefore resolved from the current working directory.
func PathOpen(path string, openmode int) (int, string, error) {
	if path == "" {
		return -1, "", EINVAL
	}
	dirFd, dirName, pathPtr, pathLen := preparePath(path)

	var oflags __wasip1_oflags_t
	if (openmode & O_CREATE) != 0 {
		oflags |= OFLAG_CREATE
	}
	if (openmode & O_TRUNC) != 0 {
		oflags |= OFLAG_TRUNC
	}
	if (openmode & O_EXCL) != 0 {
		oflags |= OFLAG_EXCL
	}

	// Remove when https://github.com/bytecodealliance/wasmtime/pull/4967 is merged.
	var st Stat_t
	if errno := __wasip1_path_filestat_get(
		dirFd,
		LOOKUP_SYMLINK_FOLLOW,
		pathPtr,
		pathLen,
		&st,
	); errno != 0 && errno != ENOENT {
		return -1, "", errnoErr(errno)
	}
	if st.Filetype == FILETYPE_DIRECTORY {
		oflags |= OFLAG_DIRECTORY
		// WASM runtimes appear to return EINVAL when passing invalid
		// combination of flags to open directories; however, TestOpenError
		// in the os package expects EISDIR, so we precheck this condition
		// here to emulate the expected behavior.
		const invalidFlags = O_WRONLY | O_RDWR | O_CREATE | O_APPEND | O_TRUNC | O_EXCL
		if (openmode & invalidFlags) != 0 {
			return 0, "", EISDIR
		}
	}

	var rights __wasip1_rights_t
	switch openmode & (O_RDONLY | O_WRONLY | O_RDWR) {
	case O_RDONLY:
		rights = fullRights & ^writeRights
	case O_WRONLY:
		rights = fullRights & ^readRights
	case O_RDWR:
		rights = fullRights
	}

	var fdflags __wasip1_fdflags_t
	if (openmode & O_APPEND) != 0 {
		fdflags |= FDFLAG_APPEND
	}
	if (openmode & O_SYNC) != 0 {
		fdflags |= FDFLAG_SYNC
	}

	var fd __wasip1_fd_t
	errno := __wasip1_path_open(
		dirFd,
		LOOKUP_SYMLINK_FOLLOW,
		pathPtr,
		pathLen,
		oflags,
		rights,
		fullRights,
		fdflags,
		&fd,
	)
	if errno != 0 {
		return -1, "", errnoErr(errno)
	}
	path = joinPath(dirName, unsafe.String(pathPtr, pathLen))
	return int(fd), path, errnoErr(errno)
}

func Close(fd int) error {
	errno := __wasip1_fd_close(__wasip1_fd_t(fd))
	return errnoErr(errno)
}

func CloseOnExec(fd int) {
	// nothing to do - no exec
}

func Mkdir(path string, perm uint32) error {
	if path == "" {
		return EINVAL
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_create_directory(dirFd, pathPtr, pathLen)
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
		return EINVAL
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_filestat_get(dirFd, LOOKUP_SYMLINK_FOLLOW, pathPtr, pathLen, st)
	setDefaultMode(st)
	return errnoErr(errno)
}

func Lstat(path string, st *Stat_t) error {
	if path == "" {
		return EINVAL
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_filestat_get(dirFd, 0, pathPtr, pathLen, st)
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
		return EINVAL
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_unlink_file(dirFd, pathPtr, pathLen)
	return errnoErr(errno)
}

func Rmdir(path string) error {
	if path == "" {
		return EINVAL
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_remove_directory(dirFd, pathPtr, pathLen)
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
	return ENOSYS
}

func Fchown(fd int, uid, gid int) error {
	return ENOSYS
}

func Lchown(path string, uid, gid int) error {
	return ENOSYS
}

func UtimesNano(path string, ts []Timespec) error {
	if path == "" {
		return EINVAL
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_filestat_set_times(
		dirFd,
		LOOKUP_SYMLINK_FOLLOW,
		pathPtr,
		pathLen,
		__wasip1_timestamp_t(TimespecToNsec(ts[0])),
		__wasip1_timestamp_t(TimespecToNsec(ts[1])),
		FILESTAT_SET_ATIM|FILESTAT_SET_MTIM,
	)
	return errnoErr(errno)
}

func Rename(from, to string) error {
	if from == "" || to == "" {
		return EINVAL
	}
	oldDirFd, _, oldPathPtr, oldPathLen := preparePath(from)
	newDirFd, _, newPathPtr, newPathLen := preparePath(to)
	errno := __wasip1_path_rename(
		oldDirFd,
		oldPathPtr,
		oldPathLen,
		newDirFd,
		newPathPtr,
		newPathLen,
	)
	return errnoErr(errno)
}

func Truncate(path string, length int64) error {
	if path == "" {
		return EINVAL
	}
	fd, err := Open(path, O_WRONLY, 0)
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
	return cwd, nil
}

func Chdir(path string) error {
	if path == "" {
		return EINVAL
	}

	dir := "/"
	if !isAbs(path) {
		dir = cwd
	}
	path = joinPath(dir, path)

	var stat Stat_t
	dirFd, _, pathPtr, pathLen := preparePath(path)
	errno := __wasip1_path_filestat_get(dirFd, LOOKUP_SYMLINK_FOLLOW, pathPtr, pathLen, &stat)
	if errno != 0 {
		return errnoErr(errno)
	}
	if stat.Filetype != FILETYPE_DIRECTORY {
		return ENOTDIR
	}
	cwd = path
	return nil
}

func Readlink(path string, buf []byte) (n int, err error) {
	if path == "" {
		return 0, EINVAL
	}
	if len(buf) == 0 {
		return 0, nil
	}
	dirFd, _, pathPtr, pathLen := preparePath(path)
	var bufused size_t
	errno := __wasip1_path_readlink(
		dirFd,
		pathPtr,
		pathLen,
		&buf[0],
		size_t(len(buf)),
		&bufused,
	)
	// For some reason wasmtime returns ERANGE when the output buffer is
	// shorter than the symbolic link value. os.Readlink expects a nil
	// error and uses the fact that n is greater or equal to the buffer
	// length to assume that it needs to try again with a larger size.
	// This condition is handled in os.Readlink.
	return int(bufused), errnoErr(errno)
}

func Link(path, link string) error {
	if path == "" || link == "" {
		return EINVAL
	}
	oldDirFd, _, oldPathPtr, oldPathLen := preparePath(path)
	newDirFd, _, newPathPtr, newPathLen := preparePath(link)
	errno := __wasip1_path_link(
		oldDirFd,
		0,
		oldPathPtr,
		oldPathLen,
		newDirFd,
		newPathPtr,
		newPathLen,
	)
	return errnoErr(errno)
}

func Symlink(path, link string) error {
	if path == "" || link == "" {
		return EINVAL
	}
	dirFd, _, pathPtr, pathlen := preparePath(link)
	errno := __wasip1_path_symlink(
		unsafe.StringData(path),
		size_t(len(path)),
		dirFd,
		pathPtr,
		pathlen,
	)
	return errnoErr(errno)
}

func Fsync(fd int) error {
	errno := __wasip1_fd_sync(__wasip1_fd_t(fd))
	return errnoErr(errno)
}

func makeIOVec(b []byte) *__wasip1_iovec_t {
	return &__wasip1_iovec_t{
		buf:     uintptr_t(uintptr(unsafe.Pointer(unsafe.SliceData(b)))),
		buf_len: size_t(len(b)),
	}
}

func Read(fd int, b []byte) (int, error) {
	var nread size_t
	errno := __wasip1_fd_read(__wasip1_fd_t(fd), makeIOVec(b), 1, &nread)
	runtime.KeepAlive(b)
	return int(nread), errnoErr(errno)
}

func Write(fd int, b []byte) (int, error) {
	var nwritten size_t
	errno := __wasip1_fd_write(__wasip1_fd_t(fd), makeIOVec(b), 1, &nwritten)
	runtime.KeepAlive(b)
	return int(nwritten), errnoErr(errno)
}

func Pread(fd int, b []byte, offset int64) (int, error) {
	var nread size_t
	errno := __wasip1_fd_pread(__wasip1_fd_t(fd), makeIOVec(b), 1, __wasip1_filesize_t(offset), &nread)
	runtime.KeepAlive(b)
	return int(nread), errnoErr(errno)
}

func Pwrite(fd int, b []byte, offset int64) (int, error) {
	var nwritten size_t
	errno := __wasip1_fd_pwrite(__wasip1_fd_t(fd), makeIOVec(b), 1, __wasip1_filesize_t(offset), &nwritten)
	runtime.KeepAlive(b)
	return int(nwritten), errnoErr(errno)
}

func Seek(fd int, offset int64, whence int) (int64, error) {
	var newoffset __wasip1_filesize_t
	errno := __wasip1_fd_seek(__wasip1_fd_t(fd), __wasip1_filedelta_t(offset), __wasip1_whence_t(whence), &newoffset)
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
	errno := __wasip1_random_get(unsafe.SliceData(b), size_t(len(b)))
	return errnoErr(errno)
}
