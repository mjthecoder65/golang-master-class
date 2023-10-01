package main

// logging in go, using the log package.
// Go standard library provides log package as basic infrastructure for logging.
// The main purposes of loggins is to get a trace of what is happening in programs.
// Logging can be providing code tracing, profiling and analytics.
// Logging is a way to find those bugs and learn more about how the program is functioning.
// import log package to use logging.

import (
	"log"
	"net/smtp"
)

// In its simplest form, it format messages and send send them to Standard Error.

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

// Try to connect to SMPT server that does not exist
// The program will print the error message and exit.

func TestingLogging() {
	client, err := smtp.Dial("smtp.smail.com:25")

	if err != nil {
		log.Fatalln(err)
	}

	client.Data()
}
