// Package leap determines whether a given year is a leap year.
package leap

// Function IsLeapYear returns true if year is a leap year and false if not.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
