// Code generated by linux/mkall.go generatePtraceRegSet("loong64"). DO NOT EDIT.

package unix

import "unsafe"

// PtraceGetRegSetLoong64 fetches the registers used by loong64 binaries.
func PtraceGetRegSetLoong64(pid, addr int, regsout *PtraceRegsLoong64) error {
	iovec := Iovec{(*byte)(unsafe.Pointer(regsout)), uint64(unsafe.Sizeof(*regsout))}
	return ptrace(PTRACE_GETREGSET, pid, uintptr(addr), uintptr(unsafe.Pointer(&iovec)))
}

// PtraceSetRegSetLoong64 sets the registers used by loong64 binaries.
func PtraceSetRegSetLoong64(pid, addr int, regs *PtraceRegsLoong64) error {
	iovec := Iovec{(*byte)(unsafe.Pointer(regs)), uint64(unsafe.Sizeof(*regs))}
	return ptrace(PTRACE_SETREGSET, pid, uintptr(addr), uintptr(unsafe.Pointer(&iovec)))
}

// PtraceRegsLoongarch is the registers used by loong binaries.
type PtraceRegsLoongarch struct {
	Uregs [18]uint32
}

// PtraceGetRegsLoongarch fetches the registers used by loong binaries.
func PtraceGetRegsLoongarch(pid int, regsout *PtraceRegsLoongarch) error {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsLoongarch sets the registers used by loong binaries.
func PtraceSetRegsLoongarch(pid int, regs *PtraceRegsLoongarch) error {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

// PtraceRegsLoong64 is the registers used by loong64 binaries.
type PtraceRegsLoong64 struct {
	Regs   [31]uint64
	Sp     uint64
	Pc     uint64
	Pstate uint64
}

// PtraceGetRegsLoong64 fetches the registers used by loong64 binaries.
func PtraceGetRegsLoong64(pid int, regsout *PtraceRegsLoong64) error {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsLoong64 sets the registers used by loong64 binaries.
func PtraceSetRegsLoong64(pid int, regs *PtraceRegsLoong64) error {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}
