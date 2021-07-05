package main

import (
  "fmt"
	"log"
	"rsc.io/quote"
	"example.com/greetings"
	"cmon.consulting/go/greetings"
)


func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}

