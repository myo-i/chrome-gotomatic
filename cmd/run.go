package cmd

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "build simple container",
	Run: func(cmd *cobra.Command, args []string) {
		switch os.Args[1] {
		case "run":
			run()
		default:
			panic("Error")
		}
	},
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	err := cmd.Run()
	if err != nil {
		panic("Error")
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}

type SysProcAttr struct {
	HideWindow                 bool
	CmdLine                    string // used if non-empty, else the windows command line is built by escaping the arguments passed to StartProcess
	CreationFlags              uint32
	Token                      Token               // if set, runs new process in the security context represented by the token
	ProcessAttributes          *SecurityAttributes // if set, applies these security attributes as the descriptor for the new process
	ThreadAttributes           *SecurityAttributes // if set, applies these security attributes as the descriptor for the main thread of the new process
	NoInheritHandles           bool                // if set, each inheritable handle in the calling process is not inherited by the new process
	AdditionalInheritedHandles []Handle            // a list of additional handles, already marked as inheritable, that will be inherited by the new process
	ParentProcess              Handle              // if non-zero, the new process regards the process given by this handle as its parent process, and AdditionalInheritedHandles, if set, should exist in this parent process
}

// type SysProcAttr struct {
// 	Chroot     string      // Chroot.
// 	Credential *Credential // Credential.
// 	Ptrace     bool        // Enable tracing.
// 	Setsid     bool        // Create session.
// 	// Setpgid sets the process group ID of the child to Pgid,
// 	// or, if Pgid == 0, to the new child's process ID.
// 	Setpgid bool
// 	// Setctty sets the controlling terminal of the child to
// 	// file descriptor Ctty. Ctty must be a descriptor number
// 	// in the child process: an index into ProcAttr.Files.
// 	// This is only meaningful if Setsid is true.
// 	Setctty bool
// 	Noctty  bool // Detach fd 0 from controlling terminal
// 	Ctty    int  // Controlling TTY fd
// 	// Foreground places the child process group in the foreground.
// 	// This implies Setpgid. The Ctty field must be set to
// 	// the descriptor of the controlling TTY.
// 	// Unlike Setctty, in this case Ctty must be a descriptor
// 	// number in the parent process.
// 	Foreground bool
// 	Pgid       int // Child's process group ID if Setpgid.
// }
