// Go exercise rail-fence-cipher.
package railfence

// Encode a text using the rail fence cipher.
func Encode(plainText string, rails int) string {
	plainTextRunes := []rune(plainText)
	cipherText := make([]rune, len(plainTextRunes))
	encodeMap := makeEncodeMap(len(plainTextRunes), rails)
	for i := 0; i < len(plainTextRunes); i++ {
		cipherText[i] = plainTextRunes[encodeMap[i]]
	}
	return string(cipherText)
}

// Decode a text using the rail fence cipher.
func Decode(cipherText string, rails int) string {
	cipherTextRunes := []rune(cipherText)
	plainText := make([]rune, len(cipherTextRunes))
	decodeMap := makeDecodeMap(len(cipherTextRunes), rails)
	for i := 0; i < len(cipherTextRunes); i++ {
		plainText[i] = cipherTextRunes[decodeMap[i]]
	}
	return string(plainText)
}

// Make the message encode map: encodeMap[i] = j means the i'th rune in the cipher text is the j'th
// rune in the plain text.
func makeEncodeMap(msgLen, rails int) []int {
	railMaps := make([][]int, rails)
	rail := 0
	direction := 1
	for i := 0; i < msgLen; i++ {
		railMaps[rail] = append(railMaps[rail], i)
		rail, direction = moveRail(rail, direction, rails)
	}
	encodeMap := make([]int, 0, msgLen)
	for rail = 0; rail < rails; rail++ {
		encodeMap = append(encodeMap, railMaps[rail]...)
	}
	return encodeMap
}

// Make the message decode map: encodeMap[i] = j means the i'th rune in the plain text is the j'th
// rune in the cipher text.
func makeDecodeMap(msgLen, rails int) []int {
	encodeMap := makeEncodeMap(msgLen, rails)
	decodeMap := make([]int, msgLen)
	for i := 0; i < msgLen; i++ {
		decodeMap[encodeMap[i]] = i
	}
	return decodeMap
}

// Decide what the next rail and the new direction is for the encoded message given a current rail,
// a current direction, and a number of rails.
func moveRail(rail, direction, rails int) (int, int) {
	switch {
	case rails == 1:
		return rail, direction
	case direction == 1:
		if rail == rails-1 {
			return rail - 1, -1
		}
		return rail + 1, 1
	case direction == -1:
		if rail == 0 {
			return 1, 1
		}
		return rail - 1, -1
	}
	panic("Invalid direction")
}
