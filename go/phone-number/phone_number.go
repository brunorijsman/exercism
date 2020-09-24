// Exercism exercise phone-number.
package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Normalize a phone number into digits only.
func Number(phone string) (string, error) {
	// Remove spaces, braces, dashes, and dots.
	for _, removeString := range []string{" ", "(", ")", "-", "."} {
		phone = strings.ReplaceAll(phone, removeString, "")
	}
	// Remove leading +1 or 1
	if phone[0:2] == "+1" {
		phone = phone[2:]
	} else if phone[0:1] == "1" {
		phone = phone[1:]
	}
	// Length must be 10.
	if len(phone) != 10 {
		return "", errors.New("Length must be 10")
	}
	// Must be all digits.
	for _, r := range phone {
		if !unicode.IsDigit(r) {
			return "", errors.New("Must be all digits")
		}
	}
	// Area code must not start with 0 or 1.
	if phone[0] == '0' || phone[0] == '1' {
		return "", errors.New("Area code must not start with 0 or 1")
	}
	// Exchange code must not start with 0 or 1.
	if phone[3] == '0' || phone[3] == '1' {
		return "", errors.New("Exchange code must not start with 0 or 1")
	}
	return phone, nil
}

// Extract the area code from a phone number.
func AreaCode(phone string) (string, error) {
	phone, err := Number(phone)
	if err != nil {
		return "", err
	}
	return phone[0:3], nil
}

// Normalize a phone number to (xxx) xxx-xxxxx.
func Format(phone string) (string, error) {
	phone, err := Number(phone)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phone[0:3], phone[3:6], phone[6:10]), nil
}
