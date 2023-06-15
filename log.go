package main

import (
	"log"
	"os"
)

// creating a set of varibles that will acts as logger
var (
	SessionTracker *log.Logger
	InfoLogger     *log.Logger
	ErrorLoffer    *log.Logger
	warningLogger  *log.Logger
)

// init functions is an another type of function in go that will execute before main function
func init() {
	file, err := os.OpenFile("log_history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	SessionTracker = log.New(file, "Session:", log.Ldate|log.Lshortfile|log.Ltime) // setoutput , alias , flags
	InfoLogger = log.New(file, "Info:", log.Ldate|log.Lshortfile|log.Ltime)
	ErrorLoffer = log.New(file, "Error:", log.Ldate|log.Lshortfile|log.Ltime)
	warningLogger = log.New(file, "warning:", log.Ldate|log.Lshortfile|log.Ltime)
}

func main() {
	SessionTracker.Println("Session has started")

	InfoLogger.Println("This is some info")
	warningLogger.Println("This is some warning")
	ErrorLoffer.Println("This is some error")

	SessionTracker.Println("Session has ended")
	SessionTracker.Println("___________________________")
}
