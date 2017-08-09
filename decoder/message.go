package decoder

import (
	"bytes"
)

var Message bytes.Buffer

/**
 * Append a message to existing message
 * @function AppendMessage
 * @param message String The message
 */
func AppendMessage(message string) {
	for i := 0; i < len(message); i++ {
		Message.WriteString(string([]rune(message)[i]))
	}
}

/**
 * Clears the existing message
 * @function ClearMessage
 */
func ClearMessage() {
	Message.Reset()
}

/**
 * Check if message is a header or message
 * @function IsMessage
 * @returns bool
 */
func IsMessage(message string) bool {
	onlyZerosAndOnes := true
	for i := 0; i < len(message); i++ {
		char := []rune(message)[i]
		if char != '1' && char != '0' {
			onlyZerosAndOnes = false
			break
		}
	}
	return onlyZerosAndOnes
}
