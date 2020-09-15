package reverse

func Reverse(str string) string {
	runesStr := []rune(str)
	nrRunes := len(runesStr)
	reversedRunesStr := make([]rune, nrRunes)
	for i, r := range runesStr {
		reversedRunesStr[nrRunes-i-1] = r
	}
	return string(reversedRunesStr)
}
