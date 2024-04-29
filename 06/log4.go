package main

import (
	"log"
	"os"
)

func main() {
	fileName := "New.log"
	defer os.Remove(fileName)
	logFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("open file error")
	}
	defer logFile.Close()

	debugLog := log.New(logFile, "[Info]", log.Llongfile)
	debugLog.Println("Info Level Message")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("Debug Level Message")
}
