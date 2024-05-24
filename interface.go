package main

import (
	"fmt"
)

type message interface {
	getMessage() string
}

type BirthdayMessage struct {
	birthdayTime  string
	recipientName string
}

// Define the getMessage method for BirthdayMessage
func (sr BirthdayMessage) getMessage() {
 fmt.Printf("Welcome to the birthday party at %s, %s!\n", sr.birthdayTime, sr.recipientName)
}

func main() {
	// Initialize an instance of BirthdayMessage
	pm := BirthdayMessage{
		birthdayTime:  "9 PM",
		recipientName: "Jyo",
	}

	// Call the getMessage method on the pm instance and print the result
	 pm.getMessage()
	 
}
