// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package runtime

import (
	"runtime/internal/atomic"
	"unsafe"
)

type uintptr_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-size-u32
type size_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-errno-variant
type __wasi_errno_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-filesize-u64
type __wasi_filesize_t uint64

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-timestamp-u64
type __wasi_timestamp_t uint64

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd-handle
type __wasi_fd_t uint32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-clockid-variant
type __wasi_clockid_t uint32

const (
	__WASI_CLOCK_REALTIME           __wasi_clockid_t = 0
	__WASI_CLOCK_MONOTONIC          __wasi_clockid_t = 1
	__WASI_CLOCK_PROCESS_CPUTIME_ID __wasi_clockid_t = 2
	__WASI_CLOCK_THREAD_CPUTIME_ID  __wasi_clockid_t = 3
)

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-iovec-record
type __wasi_iovec_t struct {
	buf     uintptr_t
	buf_len size_t
}

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-ciovec-record
type __wasi_ciovec_t struct {
	buf     uintptr_t
	buf_len size_t
}

//go:wasmimport wasi_snapshot_preview1 args_get
//go:noescape
func __wasi_args_get(
	argv *uintptr_t,
	argv_buf *byte,
) __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 args_sizes_get
//go:noescape
func __wasi_args_sizes_get(
	argc *size_t,
	argv_buf_size *size_t,
) __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 clock_time_get
//go:noescape
func __wasi_clock_time_get(
	clock_id __wasi_clockid_t,
	precision __wasi_timestamp_t,
	time *__wasi_timestamp_t,
) __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 environ_get
//go:noescape
func __wasi_environ_get(
	environ *uintptr_t,
	environ_buf *byte,
) __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 environ_sizes_get
//go:noescape
func __wasi_environ_sizes_get(
	environ_count *size_t,
	environ_buf_size *size_t,
) __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 proc_exit
func __wasi_proc_exit(
	code int32,
)

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func __wasi_fd_write(
	fd __wasi_fd_t,
	iovs *__wasi_ciovec_t,
	iovs_len size_t,
	nwritten *size_t,
) __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 sched_yield
func __wasi_sched_yield() __wasi_errno_t

//go:wasmimport wasi_snapshot_preview1 random_get
//go:noescape
func __wasi_random_get(
	buf *byte,
	buf_len size_t,
) __wasi_errno_t

type __wasi_eventtype_t uint8

const (
	__wasi_eventtype_clock __wasi_eventtype_t = iota
	__wasi_eventtype_fd_read
	__wasi_eventtype_fd_write
)

type __wasi_eventrwflags_t uint16

const (
	__wasi_fd_readwrite_hangup __wasi_eventrwflags_t = 1 << iota
)

type __wasi_userdata_t uint64

type __wasi_event_t struct {
	userdata    __wasi_userdata_t
	error       uint16 // TODO: this should be __wasi_errno_t but the compiler rejects uint16 as argument to imported functions
	typ         __wasi_eventtype_t
	fdReadwrite __wasi_event_fd_readwrite_t
}

type __wasi_event_fd_readwrite_t struct {
	nbytes __wasi_filesize_t
	flags  __wasi_eventrwflags_t
}

type __wasi_subclockflags_t uint16

const (
	__wasi_subscription_clock_abstime __wasi_subclockflags_t = 1 << iota
)

type __wasi_subscription_clock_t struct {
	id        __wasi_clockid_t
	timeout   __wasi_timestamp_t
	precision __wasi_timestamp_t
	flags     __wasi_subclockflags_t
}

type __wasi_subscription_t struct {
	userdata __wasi_userdata_t
	u        __wasi_subscription_u
}

type __wasi_subscription_u [5]uint64

func (u *__wasi_subscription_u) __wasi_eventtype_t() *__wasi_eventtype_t {
	return (*__wasi_eventtype_t)(unsafe.Pointer(&u[0]))
}

func (u *__wasi_subscription_u) __wasi_subscription_clock_t() *__wasi_subscription_clock_t {
	return (*__wasi_subscription_clock_t)(unsafe.Pointer(&u[1]))
}

//go:wasmimport wasi_snapshot_preview1 poll_oneoff
//go:noescape
func __wasi_poll_oneoff(
	in *__wasi_subscription_t,
	out *__wasi_event_t,
	nsubscriptions size_t,
	nevents *size_t,
) __wasi_errno_t

func exit(code int32) {
	__wasi_proc_exit(code)
}

func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	if fd > 2 {
		throw("runtime.write to fd > 2 is unsupported")
	}
	iov := __wasi_ciovec_t{
		buf:     uintptr_t(uintptr(p)),
		buf_len: size_t(n),
	}
	var nwritten size_t
	if __wasi_fd_write(__wasi_fd_t(fd), &iov, 1, &nwritten) != 0 {
		throw("__wasi_fd_write failed")
	}
	return int32(nwritten)
}

// Stubs so tests can link correctly. These should never be called.
func open(name *byte, mode, perm int32) int32        { panic("not implemented") }
func closefd(fd int32) int32                         { panic("not implemented") }
func read(fd int32, p unsafe.Pointer, n int32) int32 { panic("not implemented") }

func usleep(usec uint32) {
	var in __wasi_subscription_t
	var out __wasi_event_t
	var nevents size_t

	eventtype := in.u.__wasi_eventtype_t()
	*eventtype = __wasi_eventtype_clock

	subscription := in.u.__wasi_subscription_clock_t()
	subscription.id = __WASI_CLOCK_MONOTONIC
	subscription.timeout = __wasi_timestamp_t(usec) * 1e3
	subscription.precision = 1e3

	if __wasi_poll_oneoff(&in, &out, 1, &nevents) != 0 {
		throw("wasi_snapshot_preview1.poll_oneoff")
	}
}

//go:nosplit
func usleep_no_g(usec uint32) {
	usleep(usec)
}

// func exitThread(wait *uint32)
// FIXME: wasm doesn't have atomic yet
func exitThread(wait *atomic.Uint32)

type mOS struct{}

func osyield()

//go:nosplit
func osyield_no_g() {
	osyield()
}

func sigpanic() {
	// FIXME
}

type sigset struct{}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	mp.gsignal = malg(32 * 1024)
	mp.gsignal.m = mp
}

//go:nosplit
func msigsave(mp *m) {
}

//go:nosplit
func sigsave(p *sigset) {
	// FIXME
}

//go:nosplit
func msigrestore(sigmask sigset) {
}

//go:nosplit
//go:nowritebarrierrec
func clearSignalHandlers() {
}

//go:nosplit
func sigblock(exiting bool) {
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, cannot allocate memory.
func minit() {
}

// Called from dropm to undo the effect of an minit.
func unminit() {
}

func mdestroy(mp *m) {
	// FIXME
}

func osinit() {
	ncpu = 1
	getg().m.procid = 2
	physPageSize = 64 * 1024
}

// wasm has no signals
const _NSIG = 0

func signame(sig uint32) string {
	return ""
}

func crash() {
	*(*int32)(nil) = 0
}

func getRandomData(r []byte) {
	if __wasi_random_get(&r[0], size_t(len(r))) != 0 {
		throw("__wasi_random_get failed")
	}
}

func goenvs() {
	// arguments
	var argc size_t
	var argv_buf_size size_t
	if __wasi_args_sizes_get(&argc, &argv_buf_size) != 0 {
		throw("__wasi_args_sizes_get failed")
	}

	argslice = make([]string, argc)
	if argc > 0 {
		argv := make([]uintptr_t, argc)
		argv_buf := make([]byte, argv_buf_size)
		if __wasi_args_get(&argv[0], &argv_buf[0]) != 0 {
			throw("__wasi_args_get failed")
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
	if __wasi_environ_sizes_get(&environ_count, &environ_buf_size) != 0 {
		throw("__wasi_environ_sizes_get failed")
	}

	envs = make([]string, environ_count)
	if environ_count > 0 {
		environ := make([]uintptr_t, environ_count)
		environ_buf := make([]byte, environ_buf_size)
		if __wasi_environ_get(&environ[0], &environ_buf[0]) != 0 {
			throw("__wasi_environ_get failed")
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
	var time __wasi_timestamp_t
	if __wasi_clock_time_get(__WASI_CLOCK_REALTIME, 0, &time) != 0 {
		throw("__wasi_clock_time_get failed")
	}
	return int64(time / 1000000000), int32(time % 1000000000)
}

func nanotime1() int64 {
	var time __wasi_timestamp_t
	if __wasi_clock_time_get(__WASI_CLOCK_MONOTONIC, 0, &time) != 0 {
		throw("__wasi_clock_time_get failed")
	}
	return int64(time)
}

func initsig(preinit bool) {
}

// May run with m.p==nil, so write barriers are not allowed.
//
//go:nowritebarrier
func newosproc(mp *m) {
	panic("newosproc: not implemented")
}

func setProcessCPUProfiler(hz int32) {}
func setThreadCPUProfiler(hz int32)  {}
func sigdisable(uint32)              {}
func sigenable(uint32)               {}
func sigignore(uint32)               {}

//go:linkname os_sigpipe os.sigpipe
func os_sigpipe() {
	throw("too many writes on closed pipe")
}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed runtime·fastrand().
	// runtime·nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	// TODO: need more entropy to better seed fastrand.
	return nanotime()
}

//go:linkname syscall_now syscall.now
func syscall_now() (sec int64, nsec int32) {
	sec, nsec, _ = time_now()
	return
}

// gsignalStack is unused on js.
type gsignalStack struct{}

const preemptMSupported = false

func preemptM(mp *m) {
	// No threads, so nothing to do.
}
