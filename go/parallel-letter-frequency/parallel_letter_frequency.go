// Go exercise parallel-letter-frequency.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Return the frequency of runes in the string s.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Use go routines to count the frequency of runes in each of the strings in strs in parallel.
func ConcurrentFrequency(strs []string) FreqMap {
	resultChan := make(chan FreqMap)
	for _, str := range strs {
		go func(str string) {
			resultChan <- Frequency(str)
		}(str)
	}
	sumFreqMap := make(FreqMap)
	for range strs {
		sumFreqMap.add(<-resultChan)
	}
	return sumFreqMap
}

// Add two frequency maps.
func (freqMap FreqMap) add(addFreqMap FreqMap) {
	for r, count := range addFreqMap {
		freqMap[r] += count
	}
}
