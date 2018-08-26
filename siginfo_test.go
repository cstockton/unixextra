package unixextra

import (
	"os/exec"
	"syscall"
	"testing"
	"unsafe"

	"golang.org/x/sys/unix"
)

func TestSiginfoRawSize(t *testing.T) {
	// Make sure the struct is 128 bytes.
	if exp, got := 128, int(unsafe.Sizeof(SiginfoRaw{})); exp != got {
		t.Fatalf(`exp size of Siginfo to be %v; got %v`, exp, got)
	}
	if exp, got := 128, int(unsafe.Sizeof(SiginfoSigchld{})); exp != got {
		t.Fatalf(`exp size of Siginfo to be %v; got %v`, exp, got)
	}
}

func TestSiginfoSigchld(t *testing.T) {
	if testing.Short() {
		t.Skip(`short flag provided`)
	}

	cmd := exec.Command("sleep", ".250")
	cmd.Start()

	var raw SiginfoRaw
	flg := syscall.WEXITED | syscall.WNOWAIT
	_, _, e := unix.Syscall6(
		unix.SYS_WAITID, 0, 0, uintptr(unsafe.Pointer(&raw)), uintptr(flg), 0, 0)
	if e != 0 {
		t.Fatalf(`exp errno 0; got %v`, e)
	}

	si := raw.Siginfo().(*SiginfoSigchld)
	if exp, got := cmd.Process.Pid, int(si.PID); exp != got {
		t.Fatalf(`exp PID %v; got PID %v`, exp, got)
	}
}
