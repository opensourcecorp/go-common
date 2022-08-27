package osc

// logging provides consistent logging objects to cut down on copy-paste across
// other packages within OpenSourceCorp

import (
	"log"
	"os"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger

	// Suppress output if external-caller tests are being run
	isTesting bool
)

func init() {
	// The usage of bitwise OR here seems to be called "bitmask flagging", since
	// the log output option needs to be an integer and ORing their named bits
	// gives you a single integer result
	debugLogger = log.New(os.Stderr, "[ DEBUG ] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "[ INFO  ] ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stderr, "[ WARN  ] ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ ERROR ] ", log.Ldate|log.Ltime)
	fatalLogger = log.New(os.Stderr, "[ FATAL ] ", log.Ldate|log.Ltime)

	if os.Getenv("OSC_IS_TESTING") == "true" {
		isTesting = true
	} else {
		isTesting = false
	}
}

// DebugLog throws debug log messages
func DebugLog(msg string, values ...any) {
	if !isTesting {
		debugLogger.Printf(msg+"\n", values...)
	}
}

// InfoLog throws info log messages
func InfoLog(msg string, values ...any) {
	if !isTesting {
		infoLogger.Printf(msg+"\n", values...)
	}
}

// WarnLog throws warning log messages
func WarnLog(msg string, values ...any) {
	if !isTesting {
		warnLogger.Printf(msg+"\n", values...)
	}
}

// ErrorLog throws error log messages
func ErrorLog(err error, msg string, values ...any) {
	if !isTesting {
		if err != nil {
			errorLogger.Println(err.Error())
		}
		errorLogger.Printf(msg+"\n", values...)
	}
}

// FatalLog throws fatal log messages, which includes an exit call. There is no
// check here for if external tests are being run, so that callers can still see
// fatal log messages.
func FatalLog(err error, msg string, values ...any) {
	if err != nil {
		fatalLogger.Println(err.Error())
	}
	fatalLogger.Printf(msg+"\n", values...)
	os.Exit(1)
}
