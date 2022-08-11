package osc

import (
	"testing"
)

var (
	sc Syscall
)

func TestSyscall(t *testing.T) {
	// Did it run successfully?
	sc = Syscall{
		CmdLine: []string{"echo", "hello"},
	}
	sc.Exec()
	if !sc.Ok {
		t.Errorf("Expected 'echo hello' to succeed")
	}

	// Did it fail if it had output at all?
	sc = Syscall{
		CmdLine:      []string{"echo", "hello"},
		ErrCheckType: "outputGTZero",
	}
	sc.Exec()
	if sc.Ok {
		t.Errorf("Expected 'echo hello' to have command output")
	}

	// Did it fail if it had a *specific* output regex match?
	sc = Syscall{
		CmdLine:                 []string{"echo", "hello", "folks"},
		ErrCheckType:            "outputGTZero",
		OutputErrorPatternMatch: "folks",
	}
	sc.Exec()
	if sc.Ok {
		t.Errorf("Expected 'echo hello folks' to have 'folks' in the captured output")
	}
}
