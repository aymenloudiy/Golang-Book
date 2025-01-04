package main

import (
	"example/greeting"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	log.SetPrefix("greetings: ")
	// the log entry prefix and a flag to disable printing the time, source file, and line number.
    log.SetFlags(0)
	// A slice of names.
	names := []string{"nemyA","Fool","Great Love","Nimueh","Aymons"}
	// Request greeting messages for the names.
	messages, err:=greeting.Hellos(names)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of messages to the console.
	fmt.Println(messages)
}