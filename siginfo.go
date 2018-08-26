package unixextra

import (
	"fmt"
	"syscall"
	"unsafe"
)

// Siginfo is the interface common to all signals.
type Siginfo interface {
	Signal() syscall.Signal
}

// Siginfo returns a Siginfo which will be a *SiginfoSigchld if Signo is SIGCHLD
// or the same *SiginfoRaw otherwise.
func (s *SiginfoRaw) Siginfo() Siginfo {
	switch syscall.Signal(s.Signo) {
	case syscall.SIGCHLD:
		si := new(SiginfoSigchld)
		siginfoCopy(unsafe.Pointer(si), s)
		return si
	default:
		return s
	}
}

func siginfoCopy(dst unsafe.Pointer, src *SiginfoRaw) {
	// 128 byte bounds is validated after code generation and in unit tests.
	copy((*(*[128]byte)(dst))[:], (*(*[128]byte)(unsafe.Pointer(src)))[:])
}

func siginfoStr(si Siginfo) string {
	if si == nil {
		return `<nil>`
	}
	return fmt.Sprintf(`%T(%v)`, si, si.Signal())
}

func (s *SiginfoRaw) String() string     { return siginfoStr(s) }
func (s *SiginfoSigchld) String() string { return siginfoStr(s) }

func (s *SiginfoRaw) Signal() syscall.Signal     { return syscall.Signal(s.Signo) }
func (s *SiginfoSigchld) Signal() syscall.Signal { return syscall.Signal(s.Signo) }
