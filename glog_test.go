package glog

import ("time"
//	"os"
		"testing"
		"fmt"
//	"path/filepath"
		)

var (
	ti time.Time = time.Date(2015, 2, 20, 12, 34, 20, 0, time.Local)
)

func TestToolFunc(t *testing.T) {
	
	dy, dc := timeToDateClock(ti)
	if dy != 20150220 || dc != 123420 {
		t.Error("timeToDateClock format error %t ", ti)
	}
	
	name := fmt.Sprintf("%s.%s.%s.%d.log.20150220-123420",
		program,
		host,
		userName,
		pid)
	
	fname := getDailyFileName("./", ti, "")
	if fname != name {
		t.Error("getDailyFileName format error %t ", fname)
	} else {
		fmt.Printf("Daily file is %s\n", fname)
	}
}

/*
// TODO
func TestRotateFile(t *testing.T) {
	dir := os.TempDir()
	fname := fmt.Sprintf("%s.%s.%s.log",
		program,
		host,
		userName,
		)
	fname = filepath.Join(dir, fname)
	//fmt.Printf("fileName is %s\n", fname)
	
	for i := 0; i < 10 ; i++ {
		outName, err := rotateFileName(fname, 10)
		if err != nil {
			t.Error("rotateFileName format error %v ", err)
		}
		fmt.Printf("rotateFileName is %s\n", outName)
	}
}

func TestFile(t *testing.T) {
	
	l, _ := NewLogger("./", RotateDaily)

	for i :=0; i < 1000000; i++ {
		l.Info("Just Test")
	}
	l.Close()
	
	file, fname, err := create(ti, "./", false)
	fmt.Println("create", file, fname, err)
	file.Close()
}

func TestNeedRotate(t * testing.T) {
	l, err := NewLogger("./")
	if err != nil {
		t.Error("NewLogger error %v ", err)
	}
	//tt := ti
	d, _ := time.ParseDuration("24h")
	t1 := ti.Add(d)
	d, _ = time.ParseDuration("1h")
	t2 := ti.Add(d)
	d, _ = time.ParseDuration("-25h")
	t3 := ti.Add(d)
	d, _ = time.ParseDuration("-23h")
	t4 := ti.Add(d)
	
	l.SetRotateDaily(120200)
	// ti  2015-02-20 12:34:20
	if l.lastRotateTime = t1 ; l.needDailyRotate(ti) {
		t.Error("needDailyRotate ", t1, ti)
	}
	if l.lastRotateTime = t2; l.needDailyRotate(ti) {
		t.Error("needDailyRotate ", t2, ti)
	}
	if l.lastRotateTime = t3; !l.needDailyRotate(ti) {
		t.Error("needDailyRotate ", t3, ti)
	}
	if l.lastRotateTime = t4; !l.needDailyRotate(ti) {
		t.Error("needDailyRotate ", t4, ti)
	}

}
*/


/*
func TestMaxFileSizeRotate(t * testing.T) {
	l, err := NewLoggerDailyRotate("./", 130000)  // HH:MM:SS
	
	if err != nil {
		t.Error("NewLogger error %v ", err)
	}
	
	for i := 0 ; i < 1000000; i++ {
		l.Info("Just for test")
		
	}
	l.Close()
}
*/

const (
	totalNum = 100
) 

func TestMaxFileSizeRotate(t * testing.T) {
	l, err := NewLoggerFileSizeRotate("./", 1, false)
	if err != nil {
		t.Error("NewLogger error %v ", err)
	}
	l.Info("Info Just for test")
	l.Warning("Warning Just for test")
	l.Error("Error Just for test")
	l.Close()
}

/*
func TestWriteLog(t * testing.T) {
	for i := 0 ; i < totalNum; i++ {
		Info("Just for test")
	}
	Close()
}
*/