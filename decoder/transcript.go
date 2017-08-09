package decoder

import (
	"fmt"
	"bytes"
	"strconv"
)

/**
 * Gets the snippet length
 * @function getSnippetLength
 * @param snippet String That part of the message to process
 * @returns int
 */
func getSnippetLength(snippet string) int {
	length, _ := strconv.ParseInt(snippet, 2, 64)
	return int(length)
}

/**
 * Checks if header and message are available for decoding
 * @function isReadyForDecoding
 * @returns bool
 */
func isReadyForDecoding() bool {
	return len(Message.String()) > 0 && len(HeaderMapping) > 0
}

/**
 * Decode the message
 * @function DecodeMessage
 */
func DecodeMessage() {

	// Check if requirements are met
	if isReadyForDecoding() {
		var snippet bytes.Buffer

		message := Message.String()
		snippetLength := 3
		gettingSnippetLength := true

		// Go through each of the characters and evaluate if it
		// is a message length or the actual message
		for i := 0; i < len(message); i++ {

			// Capture the character
			snippet.WriteString(string([]rune(message)[i]))

			// Message length match found
			if len(snippet.String()) == int(snippetLength) {

				// Check if this is just getting the snippet length or the actual message length
				if gettingSnippetLength {
					// Just the snippet length
					snippetLength = getSnippetLength(snippet.String())
					gettingSnippetLength = false
				} else {
					// This is the actual message payload
					// Print to screen as decoded
					fmt.Print(HeaderMapping[snippet.String()])

					// Check if this mapping exists in the header
					if len(HeaderMapping[snippet.String()]) == 0 {

						// Not ion the mapping, break and go back to snippet length checking
						gettingSnippetLength = true
						snippetLength = 3
					}
				}
				snippet.Reset()
			}

		}
		fmt.Println()
	}
}
