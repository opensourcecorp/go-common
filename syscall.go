package osc

import (
	"os/exec"
	"regexp"
)

// Syscall contains data necessary to build, run, and validate external calls.
// Callers can leverage the default ("zeroed") fields of this struct -- e.g. for
// a simple "run this and see if it fails" call, the struct literal need only
// have CmdLine defined.
type Syscall struct {
	CmdLine                 []string // The actual command and its args to pass to Exec()
	ErrCheckType            string   // what kind of logical check to make against Exec() returns
	OutputErrorPatternMatch string   // used to match against, if matching against a pattern in the Exec() output
	Ok                      bool     // was Exec() successful based on conditions?
}

func (sc *Syscall) Exec() {
	var cmd *exec.Cmd
	if len(sc.CmdLine) == 1 {
		cmd = exec.Command(sc.CmdLine[0])
	} else if len(sc.CmdLine) > 1 {
		cmd = exec.Command(sc.CmdLine[0], sc.CmdLine[1:]...)
	} else {
		FatalLog(nil, "how tf u gonna give me a zero-length command")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		ErrorLog(err, "Output below:\n"+string(output))
		sc.Ok = false
		return
	}

	switch sc.ErrCheckType {
	case "nonZeroExit", "":
		// this should have failed early above, so we can just return true early
		// here
		sc.Ok = true
		return
	case "outputGTZero":
		if len(output) > 0 {
			if sc.OutputErrorPatternMatch == "" {
				ErrorLog(nil, "Output below:\n"+string(output))
				sc.Ok = false
				return
			} else {
				regex := regexp.MustCompile(sc.OutputErrorPatternMatch)
				if regex.MatchString(string(output)) {
					ErrorLog(nil, "Output below:\n"+string(output))
					sc.Ok = false
					return
				}
			}
		}
	default:
		// If it was a nonzero exit syscall, they should never get here anyway
		FatalLog(nil, "Unhandled syscall() errCheckType '%s'", sc.ErrCheckType)
	}

	sc.Ok = true
}
