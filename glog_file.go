// Go support for leveled logs, analogous to https://code.google.com/p/google-glog/
//
// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// File I/O for logs.

package glog

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// MaxSize is the maximum size of a log file in bytes.
var MaxSize uint64 = 1024 * 1024 * 1800

// logDirs lists the candidate directories for new log files.
var logDirs []string

// If non-empty, overrides the choice of directory in which to write logs.
// See createLogDirs for the full list of possible destinations.
var logDir = flag.String("log_dir", "", "If non-empty, write log files in this directory")

//
func createLogDirs(ldir string) {
	if *logDir != "" {
		logDirs = append(logDirs, *logDir)
	} else if(ldir != "") {
		logDirs = append(logDirs, ldir)
	} else {
		const default_dir = "./"
		logDirs = append(logDirs, default_dir)	
	}
	// len(logDirs) == 1
	os.MkdirAll(logDirs[0], 0666)
	fmt.Println("log dir : ", logDirs)
}

var (
	pid      = os.Getpid()
	program  = filepath.Base(os.Args[0])
	host     = "unknownhost"
	userName = "unknownuser"
)

func init() {
	h, err := os.Hostname()
	if err == nil {
		host = shortHostname(h)
	}

	current, err := user.Current()
	if err == nil {
		userName = current.Username
	}

	// Sanitize userName since it may contain filepath separators on Windows.
	userName = strings.Replace(userName, `\`, "_", -1)
}

// shortHostname returns its argument, truncating at the first period.
// For instance, given "www.google.com" it returns "www".
func shortHostname(hostname string) string {
	if i := strings.Index(hostname, "."); i >= 0 {
		return hostname[:i]
	}
	return hostname
}

// logName returns a new log file name containing tag, with start time t, and
// the name for the symlink for tag.
func logName(tag string, t time.Time) (name, link string) {
	name = fmt.Sprintf("%s.%s.%s.log.%s.%04d%02d%02d-%02d%02d%02d.%d",
		program,
		host,
		userName,
		tag,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		pid)
	return name, program + "." + tag
}

var onceLogDirs sync.Once

// create creates a new log file and returns the file and its filename, which
// contains tag ("INFO", "FATAL", etc.) and t.  If the file is created
// successfully, create also attempts to update the symlink for that tag, ignoring
// errors.
func create(/*tag string,*/ t time.Time, ldir string, dailyRotate bool) (f *os.File, filename string, err error) {
	onceLogDirs.Do(func () {
		createLogDirs(ldir)
	})
	if len(logDirs) == 0 {
		return nil, "", errors.New("log: no log dirs")
	}
	//name, link := logName(tag, t)
	var fname string
	if dailyRotate {
		fname = getDailyFileName(logDirs[0], t)
	} else {
		fname, err = getRotateFileName(logDirs[0])
	}
	
	
	if err != nil {
		return nil, "", fmt.Errorf("log: cannot create log: %v", err)
	}
	
	f, err = os.Create(fname)
	
	if err != nil {
		return nil, "", fmt.Errorf("log: cannot create log: %v", err)	
	}
	return f, fname, nil
	
}

func getDailyFileName(dir string, t time.Time) string {
	name := fmt.Sprintf("%s.%s.%s.%d.log.%04d%02d%02d-%02d%02d%02d",
		program,
		host,
		userName,
		pid,
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	return name
}


func getRotateFileName(dir string) (string, error) {
	fname := fmt.Sprintf("%s.%s.%s.%d.log",
		program,
		host,
		userName,
		pid)
	fname = filepath.Join(dir, fname)
	descName, err := rotateFileName(fname, 10)
	if err != nil {
		return descName, err
	}
	err = os.Rename(fname, descName)
	return fname, err
}

// 查找可以使用的文件夹
func rotateFileName(fname string, rotNum int) (string, error) {
	// 找最旧的一个或者找不存在的第一个如果返回空代表全是文件夹
	var t time.Time = time.Now()
	var oldName string
	for i := 0; i < rotNum; i++ {
		filename := fname + fmt.Sprintf(".%03d", i)
		fInfo, err := os.Lstat(filename)
		if err == nil {
			if  !fInfo.IsDir() && t.After(fInfo.ModTime()) {
				t = fInfo.ModTime()
				oldName = filename
			}
		} else {
			oldName = filename
			break
		}
	}
	if len(oldName) <= 0 {
		return oldName, &LoggerError{ fname + "000 can't be create" }
	}
	return oldName, nil
}
