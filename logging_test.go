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
	var (
		want, got                            string
		debugBuf, infoBuf, warnBuf, errorBuf bytes.Buffer
		match                                bool
		err                                  error
	)

	// DebugLog
	debugLogger.SetOutput(&debugBuf)
	want = `[ osc:DEBUG ].* debug`
	DebugLog("debug")
	got = readTestLogs(debugBuf)

	match, err = regexp.MatchString(want, got)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !match {
		t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
	}

	// InfoLog
	infoLogger.SetOutput(&infoBuf)
	want = `[ osc:INFO  ].* info`
	InfoLog("info")
	got = readTestLogs(infoBuf)

	match, err = regexp.MatchString(want, got)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !match {
		t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
	}

	// WarnLog
	warnLogger.SetOutput(&warnBuf)
	want = `[ osc:WARN  ].* warn`
	WarnLog("warn")
	got = readTestLogs(warnBuf)

	match, err = regexp.MatchString(want, got)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !match {
		t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
	}

	// ErrorLog
	errorLogger.SetOutput(&errorBuf)
	want = `[ osc:ERROR ].* error`
	ErrorLog(nil, "error")
	got = readTestLogs(errorBuf)

	match, err = regexp.MatchString(want, got)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !match {
		t.Errorf("Desired log pattern '%v' does not match log content '%v'\n", want, got)
	}

	// TODO: FatalLog
}
