// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package runtime

import (
	"unsafe"
)

type uintptr_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-size-u32
type size_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-errno-variant
type __wasip1_errno_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-filesize-u64
type __wasip1_filesize_t uint64

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-timestamp-u64
type __wasip1_timestamp_t uint64

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd-handle
type __wasip1_fd_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-clockid-variant
type __wasip1_clockid_t uint32

const (
	__WASIP1_CLOCK_REALTIME           __wasip1_clockid_t = 0
	__WASIP1_CLOCK_MONOTONIC          __wasip1_clockid_t = 1
	__WASIP1_CLOCK_PROCESS_CPUTIME_ID __wasip1_clockid_t = 2
	__WASIP1_CLOCK_THREAD_CPUTIME_ID  __wasip1_clockid_t = 3
)

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-iovec-record
type __wasip1_iovec_t struct {
	buf     uintptr_t
	buf_len size_t
}

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-ciovec-record
type __wasip1_ciovec_t struct {
	buf     uintptr_t
	buf_len size_t
}

//go:wasmimport wasi_snapshot_preview1 args_get
//go:noescape
func __wasip1_args_get(
	argv *uintptr_t,
	argv_buf *byte,
) __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 args_sizes_get
//go:noescape
func __wasip1_args_sizes_get(
	argc *size_t,
	argv_buf_size *size_t,
) __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 clock_time_get
//go:noescape
func __wasip1_clock_time_get(
	clock_id __wasip1_clockid_t,
	precision __wasip1_timestamp_t,
	time *__wasip1_timestamp_t,
) __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 environ_get
//go:noescape
func __wasip1_environ_get(
	environ *uintptr_t,
	environ_buf *byte,
) __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 environ_sizes_get
//go:noescape
func __wasip1_environ_sizes_get(
	environ_count *size_t,
	environ_buf_size *size_t,
) __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 proc_exit
func __wasip1_proc_exit(
	code int32,
)

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func __wasip1_fd_write(
	fd __wasip1_fd_t,
	iovs *__wasip1_ciovec_t,
	iovs_len size_t,
	nwritten *size_t,
) __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 sched_yield
func __wasip1_sched_yield() __wasip1_errno_t

//go:wasmimport wasi_snapshot_preview1 random_get
//go:noescape
func __wasip1_random_get(
	buf *byte,
	buf_len size_t,
) __wasip1_errno_t

type __wasip1_eventtype_t uint8

const (
	__wasip1_eventtype_clock __wasip1_eventtype_t = iota
	__wasip1_eventtype_fd_read
	__wasip1_eventtype_fd_write
)

type __wasip1_eventrwflags_t uint16

const (
	__wasip1_fd_readwrite_hangup __wasip1_eventrwflags_t = 1 << iota
)

type __wasip1_userdata_t uint64

type __wasip1_event_t struct {
	userdata    __wasip1_userdata_t
	error       uint16 // TODO: this should be __wasip1_errno_t but the compiler rejects uint16 as argument to imported functions
	typ         __wasip1_eventtype_t
	fdReadwrite __wasip1_event_fd_readwrite_t
}

type __wasip1_event_fd_readwrite_t struct {
	nbytes __wasip1_filesize_t
	flags  __wasip1_eventrwflags_t
}

type __wasip1_subclockflags_t uint16

const (
	__wasip1_subscription_clock_abstime __wasip1_subclockflags_t = 1 << iota
)

type __wasip1_subscription_clock_t struct {
	id        __wasip1_clockid_t
	timeout   __wasip1_timestamp_t
	precision __wasip1_timestamp_t
	flags     __wasip1_subclockflags_t
}

type __wasip1_subscription_t struct {
	userdata __wasip1_userdata_t
	u        __wasip1_subscription_u
}

type __wasip1_subscription_u [5]uint64

func (u *__wasip1_subscription_u) __wasip1_eventtype_t() *__wasip1_eventtype_t {
	return (*__wasip1_eventtype_t)(unsafe.Pointer(&u[0]))
}

func (u *__wasip1_subscription_u) __wasip1_subscription_clock_t() *__wasip1_subscription_clock_t {
	return (*__wasip1_subscription_clock_t)(unsafe.Pointer(&u[1]))
}

//go:wasmimport wasi_snapshot_preview1 poll_oneoff
//go:noescape
func __wasip1_poll_oneoff(
	in *__wasip1_subscription_t,
	out *__wasip1_event_t,
	nsubscriptions size_t,
	nevents *size_t,
) __wasip1_errno_t

func exit(code int32) {
	__wasip1_proc_exit(code)
}

func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	if fd > 2 {
		throw("runtime.write to fd > 2 is unsupported")
	}
	iov := __wasip1_ciovec_t{
		buf:     uintptr_t(uintptr(p)),
		buf_len: size_t(n),
	}
	var nwritten size_t
	if __wasip1_fd_write(__wasip1_fd_t(fd), &iov, 1, &nwritten) != 0 {
		throw("__wasip1_fd_write failed")
	}
	return int32(nwritten)
}

func usleep(usec uint32) {
	var in __wasip1_subscription_t
	var out __wasip1_event_t
	var nevents size_t

	eventtype := in.u.__wasip1_eventtype_t()
	*eventtype = __wasip1_eventtype_clock

	subscription := in.u.__wasip1_subscription_clock_t()
	subscription.id = __WASIP1_CLOCK_MONOTONIC
	subscription.timeout = __wasip1_timestamp_t(usec) * 1e3
	subscription.precision = 1e3

	if __wasip1_poll_oneoff(&in, &out, 1, &nevents) != 0 {
		throw("wasi_snapshot_preview1.poll_oneoff")
	}
}

func getRandomData(r []byte) {
	if __wasip1_random_get(&r[0], size_t(len(r))) != 0 {
		throw("__wasip1_random_get failed")
	}
}

func goenvs() {
	// arguments
	var argc size_t
	var argv_buf_size size_t
	if __wasip1_args_sizes_get(&argc, &argv_buf_size) != 0 {
		throw("__wasip1_args_sizes_get failed")
	}

	argslice = make([]string, argc)
	if argc > 0 {
		argv := make([]uintptr_t, argc)
		argv_buf := make([]byte, argv_buf_size)
		if __wasip1_args_get(&argv[0], &argv_buf[0]) != 0 {
			throw("__wasip1_args_get failed")
		}

		for i := range argslice {
			start := argv[i] - uintptr_t(uintptr(unsafe.Pointer(&argv_buf[0])))
			end := start
			for argv_buf[end] != 0 {
				end++
			}
			argslice[i] = string(argv_buf[start:end])
		}
	}

	// environment
	var environ_count size_t
	var environ_buf_size size_t
	if __wasip1_environ_sizes_get(&environ_count, &environ_buf_size) != 0 {
		throw("__wasip1_environ_sizes_get failed")
	}

	envs = make([]string, environ_count)
	if environ_count > 0 {
		environ := make([]uintptr_t, environ_count)
		environ_buf := make([]byte, environ_buf_size)
		if __wasip1_environ_get(&environ[0], &environ_buf[0]) != 0 {
			throw("__wasip1_environ_get failed")
		}

		for i := range envs {
			start := environ[i] - uintptr_t(uintptr(unsafe.Pointer(&environ_buf[0])))
			end := start
			for environ_buf[end] != 0 {
				end++
			}
			envs[i] = string(environ_buf[start:end])
		}
	}
}

func walltime() (sec int64, nsec int32) {
	return walltime1()
}

func walltime1() (sec int64, nsec int32) {
	var time __wasip1_timestamp_t
	if __wasip1_clock_time_get(__WASIP1_CLOCK_REALTIME, 0, &time) != 0 {
		throw("__wasip1_clock_time_get failed")
	}
	return int64(time / 1000000000), int32(time % 1000000000)
}

func nanotime1() int64 {
	var time __wasip1_timestamp_t
	if __wasip1_clock_time_get(__WASIP1_CLOCK_MONOTONIC, 0, &time) != 0 {
		throw("__wasip1_clock_time_get failed")
	}
	return int64(time)
}
