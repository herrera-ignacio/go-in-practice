package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	// LstdFlags turns both Ldate and Ltime
	// Lshortfile shows just the filename and the line number
	logger := log.New(logfile, "examole ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message")
	logger.Fatalln("This is a fatal error")
	logger.Println("This is the end of the function") // This will never get called
}
