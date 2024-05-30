package main

import "github.com/liu-hm19/pterm"

func main() {
	// Create a default box with PTerm and print a message in it.
	// The DefaultBox.Println method automatically starts, prints the message, and stops the box.
	pterm.DefaultBox.Println("Hello, World!")
}
