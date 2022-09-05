package osc

import (
	"bytes"
	"regexp"
	"testing"
)

func readTestLogs(b bytes.Buffer) string {
	return b.String()
}

// TestLogging tests each of the loggers. Since logs contain timestamps, want
// vs. got comparisons are done based on regexp vs. exact matches
func TestLogging(t *testing.T) {
	// We need to override `IsTesting` in this logging test block, because if
	// it's true, logs won't emit anything, and we need them to emit to the
	// buffers! So, that's why you see that happening, with a reset at the very
	// end
	IsTesting = false

	var (
		want, got                                      string
		match                                          bool
		err                                            error
		debugBuf, infoBuf, warnBuf, errorBuf, fatalBuf bytes.Buffer
	)

	t.Run("DebugLog", func(t *testing.T) {
		DebugLogger.SetOutput(&debugBuf)
		want = `\[ osc:DEBUG \].* debug`
		DebugLog("debug")
		got = readTestLogs(debugBuf)

		match, err = regexp.MatchString(want, got)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !match {
			t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
		}
	})

	t.Run("InfoLog", func(t *testing.T) {
		InfoLogger.SetOutput(&infoBuf)
		want = `\[ osc:INFO  \].* info`
		InfoLog("info")
		got = readTestLogs(infoBuf)

		match, err = regexp.MatchString(want, got)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !match {
			t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
		}
	})

	t.Run("WarnLog", func(t *testing.T) {
		WarnLogger.SetOutput(&warnBuf)
		want = `\[ osc:WARN  \].* warn`
		WarnLog("warn")
		got = readTestLogs(warnBuf)

		match, err = regexp.MatchString(want, got)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !match {
			t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
		}
	})

	t.Run("ErrorLog", func(t *testing.T) {
		ErrorLogger.SetOutput(&errorBuf)
		want = `\[ osc:ERROR \].* error`
		ErrorLog(nil, "error")
		got = readTestLogs(errorBuf)

		match, err = regexp.MatchString(want, got)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !match {
			t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
		}
	})

	t.Run("FatalLog", func(t *testing.T) {
		// For FatalLog, we actually need to RE-enable the IsTesting var,
		// because FatalLog will always produce output regardless of testing
		// status BUT will throw an os.Exit() unless testing is enabled
		IsTesting = true
		FatalLogger.SetOutput(&fatalBuf)
		want = `\[ osc:FATAL \].* fatal`
		FatalLog(nil, "fatal")
		got = readTestLogs(fatalBuf)

		match, err = regexp.MatchString(want, got)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if !match {
			t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
		}
	})

	// And here's the final reset
	IsTesting = true
}
