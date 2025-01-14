package crysuite

import (
	"crypto/rand"
	"log"
	"math/big"
)

// Extended Euclidean Algorithm to find modular inverse of a mod b
func egcd(a, b int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	}
	x1, y1, g := egcd(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return x, y, g
}

// Modular inverse of a modulo m
func modInverse(a, m int) int {
	x, _, g := egcd(a, m)
	if g != 1 {
		log.Fatal("Modular inverse does not exist")
	}
	return (x%m + m) % m
}

// Miller-Rabin Primality Test
func isPrime(n *big.Int, k int) bool {
	// If n is less than 2, it's not prime
	if n.Cmp(big.NewInt(2)) == -1 {
		return false
	}
	// If n is 2 or 3, it's prime
	if n.Cmp(big.NewInt(3)) <= 0 {
		return true
	}
	// If n is divisible by 2 or 3, it's not prime
	if n.Bit(0) == 0 || n.Mod(n, big.NewInt(3)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	// Find r and s such that n-1 = r*2^s
	r, s := big.NewInt(0), 0
	r.Sub(n, big.NewInt(1))
	for r.Bit(0) == 0 {
		r.Rsh(r, 1)
		s++
	}

	// Perform k rounds of testing
	for i := 0; i < k; i++ {
		a, _ := rand.Int(rand.Reader, new(big.Int).Sub(n, big.NewInt(4)))
		a.Add(a, big.NewInt(2)) // Ensure 2 <= a <= n-2

		x := new(big.Int).Exp(a, r, n)
		if x.Cmp(big.NewInt(1)) != 0 && x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			for j := 0; j < s-1; j++ {
				x.Exp(x, big.NewInt(2), n)
				if x.Cmp(big.NewInt(1)) == 0 {
					return false
				}
				if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
					break
				}
			}
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
				return false
			}
		}
	}
	return true
}

// Generate a random prime number with a specified number of bits
func generateRandomPrime(bits int) *big.Int {
	for {
		// Generate a random number of the specified bit length
		num, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), uint(bits)))
		if err != nil {
			log.Fatalf("Error generating random number: %v", err)
		}

		// Set the number to be odd
		num.SetBit(num, 0, 1)

		// If it is prime, return it
		if isPrime(num, 40) {
			return num
		}
	}
}

// Generate RSA keys (public and private) using either custom or random primes
func generateRSAKeys(p, q *big.Int) (publicKey, privateKey [2]*big.Int) {
	// Calculate n = p * q
	n := new(big.Int).Mul(p, q)

	// Calculate Euler's Totient: et = (p-1)*(q-1)
	et := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	// Choose public exponent e = 3
	e := big.NewInt(3)

	// Calculate private exponent d = invmod(e, et)
	d := new(big.Int).ModInverse(e, et)

	// Public key is [e, n], private key is [d, n]
	publicKey = [2]*big.Int{e, n}
	privateKey = [2]*big.Int{d, n}

	return
}

// RSA encryption (c = m^e mod n)
func encrypt(m *big.Int, e, n *big.Int) *big.Int {
	c := new(big.Int).Exp(m, e, n)
	return c
}

// RSA decryption (m = c^d mod n)
func decrypt(c *big.Int, d, n *big.Int) *big.Int {
	m := new(big.Int).Exp(c, d, n)
	return m
}
