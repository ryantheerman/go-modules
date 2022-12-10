package test

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Yo %v! Welcome! This message came from a module in a repo on github, believe it or not", name)
	return message
}
