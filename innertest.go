package glog

import (
	"../pptest"
	"bytes"
	"fmt"
	"time"
	"runtime"
)

// Benchmark
func InnterTest(totalnum int) {
	
	pptest.DoSingleThreadTest("Fprintf format func", func() {
		var buf bytes.Buffer
		for i := 0; i < totalnum; i++ {
			buf.Reset()
			_, _ = fmt.Fprintf(&buf, "123456")
		}
	})

	pptest.DoSingleThreadTest("Fprintf format func", func() {
		var buf bytes.Buffer
		for i := 0; i < totalnum; i++ {
			buf.Reset()
			_, _ = fmt.Fprintf(&buf, "%s %d", "123456", 12356, "123456")
		}
	})

	pptest.DoSingleThreadTest("time.Now func", func() {
		for i := 0; i < totalnum; i++ {
			_ = time.Now()
		}
	})

	pptest.DoSingleThreadTest("timeNow func", func() {
		for i := 0; i < totalnum; i++ {
			_ = timeNow()
		}
	})

	pptest.DoSingleThreadTest("formatHeader", func() {
		for i := 0; i < totalnum; i++ {
			buf := logging.formatHeader(infoLog, "inner.go", 0)
			logging.putBuffer(buf)
		}
	})

	pptest.DoSingleThreadTest("runtime.Caller", func() {
		for i := 0; i < totalnum; i++ {
			_, _, _, _ = runtime.Caller(3)
		}
	})

	//

	pptest.DoSingleThreadTest("header func", func() {
		for i := 0; i < totalnum; i++ {
			buf, _, _ := logging.header(infoLog, 0)
			logging.putBuffer(buf)
		}
	})

}
