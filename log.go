package alexa

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime)
	warn  = log.New(os.Stdout, "Warn: ", log.Ldate|log.Ltime)
	err   = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime)
	info  = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime)
)

func Trace(msg string) {
	trace.Println(msg)
}

func Warn(msg string) {
	warn.Println(msg)
}

func Error(msg string) {
	err.Panicln(msg)
}

func Info(msg string) {
	info.Println(msg)
}

func SetDebugMode() {
	trace = log.New(os.Stdout, "Trace: ", log.Ldate|log.Ltime)
	Info("Logger in debug mode")
}
