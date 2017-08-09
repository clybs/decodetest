package decoder

import (
	"fmt"
	"math"
	"strconv"
)

var HeaderMapping = make(map[string]string)

/**
 * Create a mapping for the header message
 * @function CreateHeaderMapping
 * @param header String The header message
 */
func CreateHeaderMapping(header string) {
	var charCountLimit, stringCountLimit float64
	var paddingFormat string

	headerLength := len(header)
	charCountLimit = 1
	i := 0

	// Do a loop with the header message
	for {
		if i < headerLength {

			// Get the string count limit
			stringCountLimit = math.Pow(float64(2), charCountLimit) - 1
			for j := 0; j < int(stringCountLimit) && i < headerLength; j++ {

				// Add 0 padding in the binary equivalent
				paddingFormat = "%0" + strconv.FormatInt(int64(charCountLimit), 10) + "b"
				key := fmt.Sprintf(paddingFormat, int64(j))

				// Map the key value pair
				HeaderMapping[key] = string([]rune(header)[i])
				i++
			}
			charCountLimit++
		} else {
			// Do a break once limit reached
			break
		}
	}
}
