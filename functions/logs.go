package functions

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func SetLog() {

	fmt.Println("Setting up log file as :", time.Now().UTC().Format("2006-01-02T15-04-05")+".log")
	logFileName := fmt.Sprintf("./logs/%s.log", time.Now().UTC().Format("2006-01-02T15-04-05"))

	fmt.Println("Opening file : ", logFileName)
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_SYNC|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)

}
