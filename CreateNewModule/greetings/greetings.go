package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string{
	var message string
	// Return a greeting that embeds the name in a message.
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}