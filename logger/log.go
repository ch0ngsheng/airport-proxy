package logger

import (
	"log"
	"os"
)

var l *log.Logger

func init() {
	file, err := os.Create("log.log")
	if err != nil {
		log.Fatalln("fail to create log.log")
	}

	l = log.New(file, "", log.LstdFlags)
}
func Printf(format string, v ...interface{}) {
	l.Printf(format+"\n", v...)
}
