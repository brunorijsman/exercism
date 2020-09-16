// Run-length encoding and decoding.
package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Run-length encode a string.
func RunLengthEncode(original string) string {
	var encoded string
	first := true
	var lastC rune
	run := 0
	for _, c := range original {
		if first || c == lastC {
			run += 1
			first = false
			lastC = c
		} else {
			if run == 1 {
				encoded += string(lastC)
			} else {
				encoded += strconv.Itoa(run) + string(lastC)
			}
			run = 1
			lastC = c
		}
	}
	if run == 1 {
		encoded += string(lastC)
	} else if run > 1 {
		encoded += strconv.Itoa(run) + string(lastC)
	}
	return encoded
}

// Run-length decode a string.
func RunLengthDecode(encoded string) string {
	var text string
	var nrString string
	encodedRunes := []rune(encoded)
	for i := 0; i < len(encodedRunes); i++ {
		c := encodedRunes[i]
		if unicode.IsDigit(c) {
			nrString += string(c)
		} else {
			if nrString == "" {
				text += string(c)
			} else {
				nr, err := strconv.Atoi(nrString)
				if err != nil {
					panic(fmt.Sprintf("Cannot convert %s to int", nrString))
				}
				text += strings.Repeat(string(c), nr)
				nrString = ""
			}
		}
	}
	if nrString != "" {
		panic("Missing character after run-length")
	}
	return text
}
