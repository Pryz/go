// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package syscall

import "runtime"

// TODO: generate with runtime/mknacl.sh, allow override with IRT.
const (
	sys_null                 = 1
	sys_nameservice          = 2
	sys_dup                  = 8
	sys_dup2                 = 9
	sys_open                 = 10
	sys_close                = 11
	sys_read                 = 12
	sys_write                = 13
	sys_lseek                = 14
	sys_stat                 = 16
	sys_fstat                = 17
	sys_chmod                = 18
	sys_isatty               = 19
	sys_brk                  = 20
	sys_mmap                 = 21
	sys_munmap               = 22
	sys_getdents             = 23
	sys_mprotect             = 24
	sys_list_mappings        = 25
	sys_exit                 = 30
	sys_getpid               = 31
	sys_sched_yield          = 32
	sys_sysconf              = 33
	sys_gettimeofday         = 40
	sys_clock                = 41
	sys_nanosleep            = 42
	sys_clock_getres         = 43
	sys_clock_gettime        = 44
	sys_mkdir                = 45
	sys_rmdir                = 46
	sys_chdir                = 47
	sys_getcwd               = 48
	sys_unlink               = 49
	sys_imc_makeboundsock    = 60
	sys_imc_accept           = 61
	sys_imc_connect          = 62
	sys_imc_sendmsg          = 63
	sys_imc_recvmsg          = 64
	sys_imc_mem_obj_create   = 65
	sys_imc_socketpair       = 66
	sys_mutex_create         = 70
	sys_mutex_lock           = 71
	sys_mutex_trylock        = 72
	sys_mutex_unlock         = 73
	sys_cond_create          = 74
	sys_cond_wait            = 75
	sys_cond_signal          = 76
	sys_cond_broadcast       = 77
	sys_cond_timed_wait_abs  = 79
	sys_thread_create        = 80
	sys_thread_exit          = 81
	sys_tls_init             = 82
	sys_thread_nice          = 83
	sys_tls_get              = 84
	sys_second_tls_set       = 85
	sys_second_tls_get       = 86
	sys_exception_handler    = 87
	sys_exception_stack      = 88
	sys_exception_clear_flag = 89
	sys_sem_create           = 100
	sys_sem_wait             = 101
	sys_sem_post             = 102
	sys_sem_get_value        = 103
	sys_dyncode_create       = 104
	sys_dyncode_modify       = 105
	sys_dyncode_delete       = 106
	sys_test_infoleak        = 109
	sys_test_crash           = 110
	sys_test_syscall_1       = 111
	sys_test_syscall_2       = 112
	sys_futex_wait_abs       = 120
	sys_futex_wake           = 121
	sys_pread                = 130
	sys_pwrite               = 131
	sys_truncate             = 140
	sys_lstat                = 141
	sys_link                 = 142
	sys_rename               = 143
	sys_symlink              = 144
	sys_access               = 145
	sys_readlink             = 146
	sys_utimes               = 147
	sys_get_random_bytes     = 150
)

// TODO: Auto-generate some day. (Hard-coded in binaries so not likely to change.)
const (
	E2BIG           Errno = 1
	EACCES          Errno = 2
	EADDRINUSE      Errno = 3
	EADDRNOTAVAIL   Errno = 4
	EAFNOSUPPORT    Errno = 5
	EAGAIN          Errno = 6
	EALREADY        Errno = 7
	EBADF           Errno = 8
	EBADMSG         Errno = 9
	EBUSY           Errno = 10
	ECANCELED       Errno = 11
	ECHILD          Errno = 12
	ECONNABORTED    Errno = 13
	ECONNREFUSED    Errno = 14
	ECONNRESET      Errno = 15
	EDEADLK         Errno = 16
	EDESTADDRREQ    Errno = 17
	EDOM            Errno = 18
	EDQUOT          Errno = 19
	EEXIST          Errno = 20
	EFAULT          Errno = 21
	EFBIG           Errno = 22
	EHOSTUNREACH    Errno = 23
	EIDRM           Errno = 24
	EILSEQ          Errno = 25
	EINPROGRESS     Errno = 26
	EINTR           Errno = 27
	EINVAL          Errno = 28
	EIO             Errno = 29
	EISCONN         Errno = 30
	EISDIR          Errno = 31
	ELOOP           Errno = 32
	EMFILE          Errno = 33
	EMLINK          Errno = 34
	EMSGSIZE        Errno = 35
	EMULTIHOP       Errno = 36
	ENAMETOOLONG    Errno = 37
	ENETDOWN        Errno = 38
	ENETRESET       Errno = 39
	ENETUNREACH     Errno = 40
	ENFILE          Errno = 41
	ENOBUFS         Errno = 42
	ENODEV          Errno = 43
	ENOENT          Errno = 44
	ENOEXEC         Errno = 45
	ENOLCK          Errno = 46
	ENOLINK         Errno = 47
	ENOMEM          Errno = 48
	ENOMSG          Errno = 49
	ENOPROTOOPT     Errno = 50
	ENOSPC          Errno = 51
	ENOSYS          Errno = 52
	ENOTCONN        Errno = 53
	ENOTDIR         Errno = 54
	ENOTEMPTY       Errno = 55
	ENOTRECOVERABLE Errno = 56
	ENOTSOCK        Errno = 57
	ENOTSUP         Errno = 58
	ENOTTY          Errno = 59
	ENXIO           Errno = 60
	EOVERFLOW       Errno = 61
	EOWNERDEAD      Errno = 62
	EPERM           Errno = 63
	EPIPE           Errno = 64
	EPROTO          Errno = 65
	EPROTONOSUPPORT Errno = 66
	EPROTOTYPE      Errno = 67
	ERANGE          Errno = 68
	EROFS           Errno = 69
	ESPIPE          Errno = 70
	ESRCH           Errno = 71
	ESTALE          Errno = 72
	ETIMEDOUT       Errno = 73
	ETXTBSY         Errno = 74
	EXDEV           Errno = 75
	ENOTCAPABLE     Errno = 76
	// needed by src/net/error_unix_test.go
	EOPNOTSUPP = ENOTSUP
)

// TODO: Auto-generate some day. (Hard-coded in binaries so not likely to change.)
var errorstr = [...]string{
	E2BIG:           "Argument list too long",
	EACCES:          "Permission denied",
	EADDRINUSE:      "Address already in use",
	EADDRNOTAVAIL:   "Address not available",
	EAFNOSUPPORT:    "Address family not supported by protocol family",
	EAGAIN:          "Try again",
	EALREADY:        "Socket already connected",
	EBADF:           "Bad file number",
	EBADMSG:         "Trying to read unreadable message",
	EBUSY:           "Device or resource busy",
	ECANCELED:       "Operation canceled.",
	ECHILD:          "No child processes",
	ECONNABORTED:    "Connection aborted",
	ECONNREFUSED:    "Connection refused",
	ECONNRESET:      "Connection reset by peer",
	EDEADLK:         "Deadlock condition",
	EDESTADDRREQ:    "Destination address required",
	EDOM:            "Math arg out of domain of func",
	EDQUOT:          "Quota exceeded",
	EEXIST:          "File exists",
	EFAULT:          "Bad address",
	EFBIG:           "File too large",
	EHOSTUNREACH:    "Host is unreachable",
	EIDRM:           "Identifier removed",
	EILSEQ:          "EILSEQ",
	EINPROGRESS:     "Connection already in progress",
	EINTR:           "Interrupted system call",
	EINVAL:          "Invalid argument",
	EIO:             "I/O error",
	EISCONN:         "Socket is already connected",
	EISDIR:          "Is a directory",
	ELOOP:           "Too many symbolic links",
	EMFILE:          "Too many open files",
	EMLINK:          "Too many links",
	EMSGSIZE:        "Message too long",
	EMULTIHOP:       "Multihop attempted",
	ENAMETOOLONG:    "File name too long",
	ENETDOWN:        "Network interface is not configured",
	ENETRESET:       "Network dropped connection on reset",
	ENETUNREACH:     "Network is unreachable",
	ENFILE:          "File table overflow",
	ENOBUFS:         "No buffer space available",
	ENODEV:          "No such device",
	ENOENT:          "No such file or directory",
	ENOEXEC:         "Exec format error",
	ENOLCK:          "No record locks available",
	ENOLINK:         "The link has been severed",
	ENOMEM:          "Out of memory",
	ENOMSG:          "No message of desired type",
	ENOPROTOOPT:     "Protocol not available",
	ENOSPC:          "No space left on device",
	ENOSYS:          "Not implemented on " + runtime.GOOS,
	ENOTCONN:        "Socket is not connected",
	ENOTDIR:         "Not a directory",
	ENOTEMPTY:       "Directory not empty",
	ENOTRECOVERABLE: "State not recoverable",
	ENOTSOCK:        "Socket operation on non-socket",
	ENOTSUP:         "Not supported",
	ENOTTY:          "Not a typewriter",
	ENXIO:           "No such device or address",
	EOVERFLOW:       "Value too large for defined data type",
	EOWNERDEAD:      "Owner died",
	EPERM:           "Operation not permitted",
	EPIPE:           "Broken pipe",
	EPROTO:          "Protocol error",
	EPROTONOSUPPORT: "Unknown protocol",
	EPROTOTYPE:      "Protocol wrong type for socket",
	ERANGE:          "Math result not representable",
	EROFS:           "Read-only file system",
	ESPIPE:          "Illegal seek",
	ESRCH:           "No such process",
	ESTALE:          "Stale file handle",
	ETIMEDOUT:       "Connection timed out",
	ETXTBSY:         "Text file busy",
	EXDEV:           "Cross-device link",
	ENOTCAPABLE:     "Capabilities insufficient",
}

// Do the interface allocations only once for common
// Errno values.
var (
	errEAGAIN error = EAGAIN
	errEINVAL error = EINVAL
	errENOENT error = ENOENT
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e Errno) error {
	switch e {
	case 0:
		return nil
	case EAGAIN:
		return errEAGAIN
	case EINVAL:
		return errEINVAL
	case ENOENT:
		return errENOENT
	}
	return e
}
