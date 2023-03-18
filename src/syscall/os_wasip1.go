package syscall

//go:wasmimport wasi_snapshot_preview1 proc_exit
func __wasip1_proc_exit(
	code int32,
)

func ProcExit(code int32) {
	__wasip1_proc_exit(code)
}
