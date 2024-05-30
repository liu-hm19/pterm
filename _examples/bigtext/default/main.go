package main

import (
	"github.com/liu-hm19/pterm"
	"github.com/liu-hm19/pterm/putils"
)

func main() {
	// Define the text to be rendered
	text := "PTerm"

	// Convert the text into a format suitable for PTerm
	letters := putils.LettersFromString(text)

	// Render the text using PTerm's default big text style
	pterm.DefaultBigText.WithLetters(letters).Render()
}
