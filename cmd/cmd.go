package cmd

import (
	"log"
	"os"
)

var ExitStatus int

func Init(name string) {
	log.SetPrefix(name + ": ")
	log.SetFlags(0)
}

func Errorln(v ...interface{}) {
	log.Println(v...)
	ExitStatus = 1
}

func Fatalln(v ...interface{}) {
	Errorln(v...)
	Exit()
}

func Exit() {
	os.Exit(ExitStatus)
}
