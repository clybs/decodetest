package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/clybs/decodetest/decoder"
)

/**
 * Main entry point
 * @function main
 */
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	decoder.ClearMessage()
	blankValueCount := 0

	for scanner.Scan() {
		header := scanner.Text()

		// Log blank values
		if len(header) == 0 {
			blankValueCount++
		}

		// Check if message part and that no blank entry added twice
		if decoder.IsMessage(header) && blankValueCount == 0 {
			// Append to existing header
			decoder.AppendMessage(header)
		} else {
			// Decode message if requirement is complete. i.e. header
			// and message part exists
			decoder.DecodeMessage()

			// Create the mapping for the header
			decoder.CreateHeaderMapping(header)

			// Reset message
			decoder.ClearMessage()
		}

		// Reset blank counter
		blankValueCount = 0
	}

	// Gotcha errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}
