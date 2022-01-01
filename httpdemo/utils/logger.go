package utils

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var ELog *logrus.Logger
var ALog *logrus.Logger

func init () {
	initErrorLog()
	initAccessLog()
}

func initErrorLog() {
	ELog = logrus.New()
	ELog.SetFormatter(&logrus.JSONFormatter{})
	fileName := "./error"+time.Now().Format("2006-01-02")+".log"
	file , err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	ww:=io.MultiWriter(os.Stdout, file)
	ELog.SetOutput(ww)
}

func initAccessLog() {
	ALog = logrus.New()
	ALog.SetFormatter(&logrus.JSONFormatter{})
	fileName := "./access"+time.Now().Format("2006-01-02")+".log"
	file , err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	ww:=io.MultiWriter(os.Stdout, file)
	ALog.SetOutput(ww)
}

