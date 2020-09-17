package cipher

import (
	"strings"
	"unicode"
)

// Internal representation of a cipher.
type cipher struct {
	distances []int
}

// Shift a rune by a given distance.
func shiftRune(r rune, distance int) rune {
	newOffset := int(r) - 'a' + distance
	if newOffset < 0 {
		newOffset += 26
	} else if newOffset >= 26 {
		newOffset -= 26
	}
	return rune(int('a') + newOffset)
}

// Create a new Caesar cipher.
func NewCaesar() Cipher {
	return NewShift(3)
}

// Create a new shift cipher.
func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return cipher{distances: []int{distance}}
}

// Create a new vigenere cipher.
func NewVigenere(key string) Cipher {
	distances := make([]int, len(key))
	if key == strings.Repeat("a", len(key)) {
		return nil
	}
	for i, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		distances[i] = int(r) - 'a'
	}
	return cipher{distances: distances}
}

// Encode a plain text.
func (c cipher) Encode(plain string) string {
	index := 0
	var builder strings.Builder
	for _, r := range plain {
		r = unicode.ToLower(r)
		if r < 'a' || r > 'z' {
			continue
		}
		builder.WriteRune(shiftRune(r, c.distances[index]))
		index = (index + 1) % len(c.distances)
	}
	return builder.String()
}

// Decode a cipher text.
func (c cipher) Decode(cipher string) string {
	index := 0
	var builder strings.Builder
	for _, r := range cipher {
		if r < 'a' || r > 'z' {
			panic("Invalid cipher text")
		}
		builder.WriteRune(shiftRune(r, -c.distances[index]))
		index = (index + 1) % len(c.distances)
	}
	return builder.String()
}
