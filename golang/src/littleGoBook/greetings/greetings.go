package greetings

import (
	"fmt"
	"errors"
)

// Hello() returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name is given, return error message
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a anme was received, return the name as embedded greeting
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message, nil
}

