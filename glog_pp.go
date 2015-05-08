package glog

import (	
	//"fmt"
	"sync/atomic"
	"time"
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

type RotateType int

const (
	RotateDaily RotateType = iota
	RotateSize				// MB
)

type LoggerError struct {
	info string
}

func (l *LoggerError)Error() string {
	return "LoggerError : " + l.info
}


//
func (l *LoggingT)setRotateDaily(time int) {
	l.rotateDaily = true	
	l.rotateTime = time // 4:00:00 -> 40000
}

func (l *LoggingT)setRotateFileSize(maxSize uint64) {
	l.rotateDaily = false
	l.rotateFileMaxSize = maxSize * uint64(1024 * 1024) // to (MB)
}

// create single file replace createfiles
func (l * LoggingT)createFile() error {
	now := time.Now()
	
	sb := &syncBuffer{
				logger: l,
			}
	
	if err := sb.rotateFile(now, l.rotateDaily); err != nil {
		return err
	}
	
	l.logFile = sb;
	return nil
}

func (l *LoggingT)Close() {
	l.lockAndFlushAll()
}

func (l *LoggingT)lockAndRotateFile() {
	l.mu.Lock()
	if l.logFile != nil {
		l.logFile.rotateFWriter()	
	}
	l.mu.Unlock()
}

func timeToDateClock(t time.Time) (int, int) {
	y, m, d := t.Date()
	h, min, s := t.Clock()
	return (y * 10000 + int(m) * 100 + d), (h * 10000 + min * 100 + s)
}

func (l *LoggingT)needDailyRotate(now time.Time) bool {
	if l.rotateDaily {
		day, t := timeToDateClock(now)
		last_day, _ := timeToDateClock(l.lastRotateTime)
		//fmt.Printf("now	 	 is %08d-%06d\n", day, t)
		//fmt.Printf("last_day is %08d set rotate time %v\n", last_day, l.rotateTime)
		//fmt.Printf("day > last_day %v t > l.rotateTime %v\n", day > last_day, t > l.rotateTime)
		// 当前日期大于上次日期 超过24 小时不考虑
		if day > last_day && t > l.rotateTime {
			l.lastRotateTime = now
			return true
		}
	}
	return false
}

func (l * LoggingT)needFileSizeRotate(currentSize uint64) bool {
	return (l.rotateFileMaxSize <= currentSize)
}


func (l *syncBuffer)rotateFWriter() error {
	return l.rotateFile(time.Now(), true)
}

func Close() {
	logging.lockAndFlushAll()
}


// create a new logger
// filename : log file name can't be empty
// rotate type : 1. rotate count, max file size 
// 				 2. daliy rotate
func NewLoggerFileSizeRotate(logPath string, maxSize uint64) (*LoggingT, error) {
	
	if logPath == "" {
		return nil, &LoggerError{ info: "logPath can't be empty" }
	}
	
	var l LoggingT
	l.logPath = logPath
	l.setRotateFileSize(maxSize)
	l.toStderr = false // outPut to stderr
	l.alsoToStderr = false
	l.logFileLevel = infoLog
	l.stderrThreshold = errorLog//l.logFileLevel
	l.setVState(0, nil, false)
	// thread-safe?
	go l.flushDaemon()
	return &l, nil
}

// create a new logger
// filename : log file name can't be empty
// rotate type : 1. rotate count, max file size 
// 				 2. daliy rotate
func NewLoggerDailyRotate(logPath string, t int) (*LoggingT, error) {
	
	if logPath == "" {
		return nil, &LoggerError{ info: "logPath can't be empty" }
	}
	
	var l LoggingT
	l.logPath = logPath
	l.setRotateDaily(t)
	l.toStderr = false // outPut to stderr
	l.alsoToStderr = false
	l.logFileLevel = infoLog
	l.stderrThreshold = errorLog//l.logFileLevel
	l.setVState(0, nil, false)
	// thread-safe?
	go l.flushDaemon()
	return &l, nil
}