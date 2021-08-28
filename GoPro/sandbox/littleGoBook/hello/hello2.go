package main

import (
  "fmt"
	"log"
	"rsc.io/quote"
	"example.com/greetings"
	"cmon.consulting/go/greetings"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// On no error, print the message and continue
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// On no error, print the message and continue
	fmt.Println(message)
	fmt.Println(quote.Go())
}

