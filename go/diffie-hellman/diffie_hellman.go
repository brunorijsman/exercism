// Exercism exercise diffie-helman
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var two = big.NewInt(2)

// Choose a random private key.
func PrivateKey(p *big.Int) *big.Int {
	var limit = new(big.Int).Sub(p, two)
	randomNr, err := rand.Int(rand.Reader, limit)
	if err != nil {
		panic("Could not generate a random number")
	}
	return randomNr.Add(randomNr, two)
}

// Compute the public key from the private key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// Choose a private-public key pair.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// Compute the shared secret from the local private key and the remote public key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
