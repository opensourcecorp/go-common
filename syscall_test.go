package osc

import (
	"testing"
)

var (
	sc Syscall
)

func TestSyscall(t *testing.T) {
	t.Run("Successful run", func(t *testing.T) {
		sc = Syscall{
			CmdLine: []string{"echo", "hello"},
		}
		sc.Exec()
		if !sc.Ok {
			t.Errorf("Expected 'echo hello' to succeed")
		}
	})

	t.Run("Fail if there's any output", func(t *testing.T) {
		sc = Syscall{
			CmdLine:      []string{"echo", "hello"},
			ErrCheckType: "outputGTZero",
		}
		sc.Exec()
		if sc.Ok {
			t.Errorf("Expected 'echo hello' to have command output")
		}
	})

	t.Run("Fail if output has regex match", func(t *testing.T) {
		sc = Syscall{
			CmdLine:                 []string{"echo", "hello", "folks"},
			ErrCheckType:            "outputGTZero",
			OutputErrorPatternMatch: "folks",
		}
		sc.Exec()
		if sc.Ok {
			t.Errorf("Expected 'echo hello folks' to have 'folks' in the captured output")
		}
	})

	t.Run("Output is captured for caller", func(t *testing.T) {
		sc = Syscall{
			CmdLine: []string{"echo", "hello"},
		}
		sc.Exec()
		if sc.Output != "hello" {
			t.Errorf("Expected 'echo hello' to have captured command output 'hello', but got '%s'", sc.Output)
		}
	})
}
