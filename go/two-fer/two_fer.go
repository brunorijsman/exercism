// Package twofer contains the solution of excercism problem "Two Fer".
package twofer

import "fmt"

// Function ShareWith returns string "One for you, one for me." if the name is empty, or "One for
// {name}, one for me." if name is non-empty.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
