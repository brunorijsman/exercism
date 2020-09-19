// Exercism exercise secret-handshake.
package secret

// Determine the secret handshake for a given code.
func Handshake(code uint) []string {
	var handshake []string
	if code&0b1 != 0 {
		handshake = append(handshake, "wink")
	}
	if code&0b10 != 0 {
		handshake = append(handshake, "double blink")
	}
	if code&0b100 != 0 {
		handshake = append(handshake, "close your eyes")
	}
	if code&0b1000 != 0 {
		handshake = append(handshake, "jump")
	}
	if code&0b10000 != 0 {
		reverse(handshake)
	}
	return handshake
}

// Reverse a secret handshake.
func reverse(handshake []string) {
	last := len(handshake) - 1
	for i := 0; i < len(handshake)/2; i++ {
		handshake[i], handshake[last-i] = handshake[last-i], handshake[i]
	}
}
