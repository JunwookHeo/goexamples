package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	LogOn     = false
	LogToFile = false
)

func init() {

	if LogOn == true {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
		if LogToFile == true {
			logFile, err := os.OpenFile(time.Now().Format("2006-01-02_15-04-05")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
			mw := io.MultiWriter(os.Stdout, logFile)
			log.SetOutput(mw)
		}
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	log.Println("Log Println")

	log.Println("End main")
}
