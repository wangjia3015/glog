package glog

import ("time"
		"testing"
		"fmt"
		)

func TestToolFunc(t *testing.T) {
	var ti time.Time = time.Date(2015, 2, 20, 12, 34, 20, 0, time.Local)
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