// Go exercise binary-search.
package binarysearch

// Search for key in a sorted slice of integers. Return the index of the key or -1 if not found.
func SearchInts(slice []int, key int) int {
	return recurseSearchInts(slice, key, 0)
}

func recurseSearchInts(slice []int, key int, start int) int {
	if len(slice) == 0 {
		return -1
	}
	midPos := len(slice) / 2
	midVal := slice[midPos]
	switch {
	case midVal == key:
		return start + midPos
	case midVal > key:
		return recurseSearchInts(slice[:midPos], key, start)
	default:
		return recurseSearchInts(slice[midPos+1:], key, start+midPos+1)
	}
}
