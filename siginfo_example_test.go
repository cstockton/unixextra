package unixextra_test

import (
	"fmt"
	"syscall"

	"github.com/cstockton/unixextra"
)

func ExampleSiginfo() {
	// If you are not sure if a syscall will be SIGCHLD, or you want to unpack
	// a different signal type on your own you may pass SiginfoRaw.
	raw := unixextra.SiginfoRaw{Signo: int32(syscall.SIGCHLD)}
	copy(raw.Raw[:], []byte{0x5a, 0x5a, 0x0, 0x0, 0xe8, 0x3})

	// Siginfo returns the concrete T of the SiginfoRaw based on it's Signo.
	si := raw.Siginfo().(*unixextra.SiginfoSigchld)
	fmt.Printf("Signo(%v) Errno(%v) Code(%v)\n", si.Signo, si.Errno, si.Code)
	fmt.Printf("PID(%v) UID(%v) Status(%v)\n", si.PID, si.UID, si.Status)

	// Output:
	// Signo(17) Errno(0) Code(0)
	// PID(23130) UID(1000) Status(0)
}
