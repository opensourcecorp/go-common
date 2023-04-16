package logging

import (
	"bytes"
	"log"
	"regexp"
	"testing"
)

// TestLogging tests each of the loggers. Since logs contain timestamps, want
// vs. got comparisons are done based on regexp vs. exact matches
func TestLogging(t *testing.T) {
	// We need to override `IsTesting` in this logging test block, because if
	// it's true, logs won't emit anything, and we need them to emit to the
	// buffers! So, that's why you see that happening, with a reset at the very
	// end
	IsTesting = false

	var (
		want, got string
		match     bool
		err       error
		debugBuf  bytes.Buffer
		infoBuf   bytes.Buffer
		warnBuf   bytes.Buffer
		errorBuf  bytes.Buffer
		fatalBuf  bytes.Buffer
	)

	type testTableItem struct {
		logger      *log.Logger
		want        string
		buffer      bytes.Buffer
		funcName    func(string, ...any)
		errFuncName func(error, string, ...any)
	}

	testTable := map[string]testTableItem{
		"DebugLog": {
			logger:   DebugLogger,
			want:     `\[ osc:DEBUG \].* DebugLog`,
			buffer:   debugBuf,
			funcName: DebugLog,
		},
		"InfoLog": {
			logger:   InfoLogger,
			want:     `\[ osc:INFO  \].* InfoLog`,
			buffer:   infoBuf,
			funcName: InfoLog,
		},
		"WarnLog": {
			logger:   WarnLogger,
			want:     `\[ osc:WARN  \].* WarnLog`,
			buffer:   warnBuf,
			funcName: WarnLog,
		},
		"ErrorLog": {
			logger:      ErrorLogger,
			want:        `\[ osc:ERROR \].* ErrorLog`,
			buffer:      errorBuf,
			errFuncName: ErrorLog,
		},
		"FatalLog": {
			logger:      FatalLogger,
			want:        `\[ osc:FATAL \].* FatalLog`,
			buffer:      fatalBuf,
			errFuncName: FatalLog,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			// For FatalLog, we actually need to RE-enable the IsTesting var,
			// because FatalLog will always produce output regardless of testing
			// status BUT will throw an os.Exit() unless testing is enabled
			if k == "FatalLog" {
				IsTesting = true
			}

			v.logger.SetOutput(&v.buffer)
			want = v.want

			// Need to dispatch based on signature, since some logs don't share the same
			if k == "ErrorLog" || k == "FatalLog" {
				v.errFuncName(nil, k)
			} else {
				v.funcName(k)
			}

			got = v.buffer.String()

			if k == "FatalLog" {
				IsTesting = false
			}

			match, err = regexp.MatchString(want, got)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if !match {
				t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
			}
		})
	}

	// And here's the final reset
	IsTesting = true
}
