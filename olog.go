// Package olog is own log
package olog

import (
	"log"
)

// log.SetPrefix("===>")
// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

var packageIsDebug = true

func init() {
	// log.SetPrefix("===>")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

// SetDebug 设置是否开启调试
func SetDebug(isDebug bool) {
	packageIsDebug = isDebug
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	log.SetFlags(flag)
}

// Printf 格式化输出日志
func Printf(format string, v ...interface{}) {
	if packageIsDebug {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Printf(format, v...)
	}
}

// Println 输出并换行
func Println(v ...interface{}) {
	if packageIsDebug {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println(v...)
	}
}
