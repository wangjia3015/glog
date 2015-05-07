package glog

import (	
	"sync/atomic"
)

func (l *LoggingT)Info(args ...interface{}) {
	l.print(infoLog, args...)
}
/*
// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func (l *LoggingT)InfoDepth(depth int, args ...interface{}) {
	l.printDepth(infoLog, depth, args...)
}
*/
// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *LoggingT)Infoln(args ...interface{}) {
	l.println(infoLog, args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *LoggingT)Infof(format string, args ...interface{}) {
	l.printf(infoLog, format, args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *LoggingT)Warning(args ...interface{}) {
	l.print(warningLog, args...)
}
/*
// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func (l *LoggingT)WarningDepth(depth int, args ...interface{}) {
	logging.printDepth(warningLog, depth, args...)
}
*/
// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *LoggingT)Warningln(args ...interface{}) {
	l.println(warningLog, args...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *LoggingT)Warningf(format string, args ...interface{}) {
	l.printf(warningLog, format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *LoggingT)Error(args ...interface{}) {
	l.print(errorLog, args...)
}
/*
// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func (l *LoggingT)ErrorDepth(depth int, args ...interface{}) {
	l.printDepth(errorLog, depth, args...)
}
*/
// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *LoggingT)Errorln(args ...interface{}) {
	l.println(errorLog, args...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *LoggingT)Errorf(format string, args ...interface{}) {
	l.printf(errorLog, format, args...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *LoggingT)Fatal(args ...interface{}) {
	l.print(fatalLog, args...)
}
/*
// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func (l *LoggingT)FatalDepth(depth int, args ...interface{}) {
	l.printDepth(fatalLog, depth, args...)
}
*/
// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *LoggingT)Fatalln(args ...interface{}) {
	l.println(fatalLog, args...)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *LoggingT)Fatalf(format string, args ...interface{}) {
	l.printf(fatalLog, format, args...)
}



// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *LoggingT)Exit(args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	l.print(fatalLog, args...)
}

/*
// ExitDepth acts as Exit but uses depth to determine which call frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func (l *LoggingT)ExitDepth(depth int, args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	l.printDepth(fatalLog, depth, args...)
}
*/
// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func (l *LoggingT)Exitln(args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	l.println(fatalLog, args...)
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *LoggingT)Exitf(format string, args ...interface{}) {
	atomic.StoreUint32(&fatalNoStacks, 1)
	l.printf(fatalLog, format, args...)
}

const {
	RotateDaily int = iota
	RotateSize				// MB
}

// create a new logger
// filename : log file name
// rotate type : 1. rotate count, max file size 
// 				 2. daliy rotate
func NewLogger(filename string, rotateType int) *LoggingT {
	var l LoggingT
	l.toStderr = false
	l.alsoToStderr = false
	// just for debug
	l.stderrThreshold = infoLog
	l.setVState(0, nil, false)
	// thread-safe?
	go l.flushDaemon()
	return &l
}
