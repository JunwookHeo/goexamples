package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/junwookheo/goexamples/logging/logrus/subpkg"
	log "github.com/sirupsen/logrus"
)

const (
	LogLevel  = log.DebugLevel
	LogToFile = false
)

func init() {
	log.SetReportCaller(true)

	if LogToFile == true {
		logFile, err := os.OpenFile(time.Now().Format("2006-01-02_15-04-05")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		mw := io.MultiWriter(logFile, os.Stdout)
		log.SetOutput(mw)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		CallerPrettyfier: caller(),
	})

	log.SetLevel(LogLevel)
}

func caller() func(*runtime.Frame) (function string, file string) {
	return func(f *runtime.Frame) (function string, file string) {
		p, _ := os.Getwd()

		return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, p), f.Line)
	}
}

func main() {
	subpkg.LogTest()

	log.Debug("Debug log")
	log.Info("Info log")
	log.Print("Print log")

	log.Warn("Warning log")
	log.Error("Error log")

	log.Fatal("Fatal log")

	log.Debug("Debug log")
	log.Info("Info log")
}
