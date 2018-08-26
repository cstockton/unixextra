// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

/*
Input to cgo -godefs.  See README.md
*/

// +godefs map struct_in_addr [4]byte /* in_addr */
// +godefs map struct_in6_addr [16]byte /* in6_addr */

package unix

/*
#define _LARGEFILE_SOURCE
#define _LARGEFILE64_SOURCE
#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE

#include <dirent.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <netpacket/packet.h>
#include <poll.h>
#include <signal.h>
#include <stdio.h>
#include <sys/epoll.h>
#include <sys/inotify.h>
#include <sys/ioctl.h>
#include <sys/mman.h>
#include <sys/mount.h>
#include <sys/param.h>
#include <sys/ptrace.h>
#include <sys/resource.h>
#include <sys/select.h>
#include <sys/signal.h>
#include <sys/statfs.h>
#include <sys/sysinfo.h>
#include <sys/time.h>
#include <sys/times.h>
#include <sys/timex.h>
#include <sys/un.h>
#include <sys/user.h>
#include <sys/utsname.h>
#include <sys/wait.h>
#include <linux/filter.h>
#include <linux/keyctl.h>
#include <linux/netlink.h>
#include <linux/perf_event.h>
#include <linux/rtnetlink.h>
#include <linux/icmpv6.h>
#include <asm/termbits.h>
#include <asm/ptrace.h>
#include <time.h>
#include <unistd.h>
#include <ustat.h>
#include <utime.h>
#include <linux/can.h>
#include <linux/if_alg.h>
#include <linux/fs.h>
#include <linux/vm_sockets.h>
#include <linux/random.h>
#include <linux/taskstats.h>
#include <linux/genetlink.h>

enum {
	sizeofPtr = sizeof(void*),
};
struct __siginfo_base {
    int Signo;
#if __SI_ERRNO_THEN_CODE
    int Errno;
    int Code;
#else
    int Code;
    int Errno;
#endif
#if __WORDSIZE == 64
    int pad0;
#endif
};

struct siginfo_raw {
    int Signo;
#if __SI_ERRNO_THEN_CODE
    int Errno;
    int Code;
#else
    int Code;
    int Errno;
#endif
#if __WORDSIZE == 64
    int pad0;
#endif
	unsigned char Raw[sizeof(siginfo_t)-sizeof(struct __siginfo_base)];
};


struct __siginfo_sigchld_base {
    int Signo;
#if __SI_ERRNO_THEN_CODE
    int Errno;
    int Code;
#else
    int Code;
    int Errno;
#endif
#if __WORDSIZE == 64
    int __pad0;
#endif
    pid_t PID;
    uid_t UID;
    int Status;
    clock_t Utime;
    clock_t Stime;
};

struct siginfo_sigchld {
    int Signo;
#if __SI_ERRNO_THEN_CODE
    int Errno;
    int Code;
#else
    int Code;
    int Errno;
#endif
#if __WORDSIZE == 64
    int pad0;
#endif
    pid_t PID;
    uid_t UID;
    int Status;
    clock_t Utime;
    clock_t Stime;
  	unsigned char Raw[sizeof(siginfo_t)-sizeof(struct __siginfo_sigchld_base)];
};

*/
import "C"

// C: basic types
type (
	Cshort     C.short
	Cint       C.int
	Clong      C.long
	Clong_long C.longlong
)

// C: basic sizes
const (
	SizeofPtr      = C.sizeofPtr
	SizeofShort    = C.sizeof_short
	SizeofInt      = C.sizeof_int
	SizeofLong     = C.sizeof_long
	SizeofLongLong = C.sizeof_longlong
)

// POSIX: Types from <types.h>
type (
	GID = C.gid_t
	UID = C.uid_t
	PID = C.pid_t
)

// POSIX: Sizes from <types.h>
const (
	SizeofGID = C.sizeof_gid_t
	SizeofUID = C.sizeof_uid_t
	SizeofPID = C.sizeof_pid_t
)

// POSIX: Types from <signal.h>
type (
	SiginfoRaw     C.struct_siginfo_raw
	SiginfoSigchld C.struct_siginfo_sigchld
)

// POSIX: Sizes of from <signal.h>
const (
	SizeofSiginfoRaw     = C.sizeof_struct_siginfo_raw
	SizeofSiginfoSigchld = C.sizeof_struct_siginfo_sigchld
)

// POSIX: Constants from <signal.h>
const (
	ILL_ILLOPC    = C.ILL_ILLOPC    // Illegal opcode.
	ILL_ILLOPN    = C.ILL_ILLOPN    // Illegal operand.
	ILL_ILLADR    = C.ILL_ILLADR    // Illegal addressing mode.
	ILL_ILLTRP    = C.ILL_ILLTRP    // Illegal trap.
	ILL_PRVOPC    = C.ILL_PRVOPC    // Privileged opcode.
	ILL_PRVREG    = C.ILL_PRVREG    // Privileged register.
	ILL_COPROC    = C.ILL_COPROC    // Coprocessor error.
	ILL_BADSTK    = C.ILL_BADSTK    // Internal stack error.
	FPE_INTDIV    = C.FPE_INTDIV    // Integer divide by zero.
	FPE_INTOVF    = C.FPE_INTOVF    // Integer overflow.
	FPE_FLTDIV    = C.FPE_FLTDIV    // Floating-point divide by zero.
	FPE_FLTOVF    = C.FPE_FLTOVF    // Floating-point overflow.
	FPE_FLTUND    = C.FPE_FLTUND    // Floating-point underflow.
	FPE_FLTRES    = C.FPE_FLTRES    // Floating-point inexact result.
	FPE_FLTINV    = C.FPE_FLTINV    // Invalid floating-point operation.
	FPE_FLTSUB    = C.FPE_FLTSUB    // Subscript out of range.
	SEGV_MAPERR   = C.SEGV_MAPERR   // Address not mapped to object.
	SEGV_ACCERR   = C.SEGV_ACCERR   // Invalid permissions for mapped object.
	BUS_ADRALN    = C.BUS_ADRALN    // Invalid address alignment.
	BUS_ADRERR    = C.BUS_ADRERR    // Nonexistent physical address.
	BUS_OBJERR    = C.BUS_OBJERR    // Object-specific hardware error.
	TRAP_BRKPT    = C.TRAP_BRKPT    // Process breakpoint.
	TRAP_TRACE    = C.TRAP_TRACE    // Process trace trap.
	CLD_EXITED    = C.CLD_EXITED    // Child has exited.
	CLD_KILLED    = C.CLD_KILLED    // Child has terminated abnormally and did not create a file.
	CLD_DUMPED    = C.CLD_DUMPED    // Child has terminated abnormally and created a file.
	CLD_TRAPPED   = C.CLD_TRAPPED   // Traced child has trapped.
	CLD_STOPPED   = C.CLD_STOPPED   // Child has stopped.
	CLD_CONTINUED = C.CLD_CONTINUED // Stopped child has continued.
	POLL_IN       = C.POLL_IN       // Data input available.
	POLL_OUT      = C.POLL_OUT      // Output buffers available.
	POLL_MSG      = C.POLL_MSG      // Input message available.
	POLL_ERR      = C.POLL_ERR      // I/O error.
	POLL_PRI      = C.POLL_PRI      // High priority input available.
	POLL_HUP      = C.POLL_HUP      // Device disconnected.
	SI_USER       = C.SI_USER       // Signal sent by ().
	SI_QUEUE      = C.SI_QUEUE      // Signal sent by ().
	SI_TIMER      = C.SI_TIMER      // Signal generated by expiration of a timer set by ().
	SI_ASYNCIO    = C.SI_ASYNCIO    // Signal generated by completion of an asynchronous I/O request.
	SI_MESGQ      = C.SI_MESGQ      // Signal generated by arrival of a message on an empty message queue.
)
