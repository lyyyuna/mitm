package mylog

import (
	"io"
	"log"
)

var logger *log.Logger

// SetLog set logger
func SetLog(l io.Writer) {
	logger = log.New(l, "[mitmproxy]", log.LstdFlags)
}

// Fatalf wrapper of logger's Fatalf
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v)
}

// Fatalln wrapper of logger's Fatalln
func Fatalln(v ...interface{}) {
	logger.Fatalln(v)
}

// Printf wrapper of logger's Printf
func Printf(format string, v ...interface{}) {
	logger.Printf(format, v)
}

// Println wrapper of logger's Println
func Println(v ...interface{}) {
	logger.Println(v)
}

// Panicln wrapper of logger's Panicln
func Panicln(v ...interface{}) {
	logger.Panicln(v)
}
