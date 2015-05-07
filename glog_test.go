package glog

import ("time"
	"os"
		"testing"
		"fmt"
	"path/filepath"
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
	
	fname := getDailyFileName("./", ti)
	if fname != name {
		t.Error("getDailyFileName format error %t ", fname)
	} else {
		fmt.Printf("Daily file is %s\n", fname)
	}
}


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
	/*
	l, _ := NewLogger("./", RotateDaily)

	for i :=0; i < 1000000; i++ {
		l.Info("Just Test")
	}
	l.Close()
	*/
	file, fname, err := create(ti, "./", false)
	fmt.Println("create", file, fname, err)
	file.Close()
}