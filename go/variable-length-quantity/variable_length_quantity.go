// Exercism exercise variable-length-encoding.
package variablelengthquantity

// Encode a series of 32-bit words into groups of 7-bit bytes.
func EncodeVarint(inWords []uint32) []uint8 {
	outBytes := []uint8{}
	for i := 0; i < len(inWords); i++ {
		inWord := inWords[i]
		var wordBytes [5]byte
		for j := 0; j < 5; j++ {
			wordBytes[j] = byte(inWord & 0x7f)
			inWord >>= 7
		}
		started := false
		for j := 4; j >= 0; j-- {
			if started || j == 0 || wordBytes[j] != 0 {
				if j != 0 {
					wordBytes[j] |= 0x80
				}
				outBytes = append(outBytes, wordBytes[j])
				started = true
			}
		}
	}
	return outBytes
}

// Decode a series of 7-bit bytes into a series of 32-bit words.
func DecodeVarint(inBytes []uint8) ([]uint32, error) {
	outWords := []uint32{}
	outWord := uint32(0)
	for i := 0; i < len(inBytes); i++ {
		inByte := inBytes[i]
		outWord = outWord<<7 | uint32(inByte&0x7f)
		if inByte&0x80 == 0 {
			outWords = append(outWords, outWord)
			outWord = 0
		}
	}
	return outWords, nil
}

func prependByte(byteSlice []uint8, byte uint8) []uint8 {
	return append([]uint8{byte}, byteSlice...)
}
