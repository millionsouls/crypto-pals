package crysuite

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math/big"
)

var ErrSharedSecretMismatch = errors.New("***Shared secret mismatch***")

func ModExp(a, b, m *big.Int) *big.Int {
	r := new(big.Int)
	r.Exp(a, b, m)
	return r
}

// Derive a symmetric key from one key pair
// a: public key
// b: private key
// p: modulus
func DHDeriveSymmetricKey(a *big.Int, b *big.Int, p *big.Int) ([32]byte, *big.Int, error) {
	secret := ModExp(a, b, p)

	sharedSecret := secret.Bytes()
	key := sha256.Sum256(sharedSecret)
	return key, secret, nil
}

// Derive a symmetric key from two key pairs
// a: Alice's key pair {a, A}
// b: Bob's key pair {b, B}
// p: modulus
func DHDeriveSymmetricKeys(a [2]*big.Int, b [2]*big.Int, p *big.Int) ([32]byte, error) {
	sAlice := ModExp(b[1], a[0], p)
	sBob := ModExp(a[1], b[0], p)

	if sAlice.Cmp(sBob) != 0 {
		return [32]byte{}, ErrSharedSecretMismatch
	}

	sharedSecret := sAlice.Bytes()
	key := sha256.Sum256(sharedSecret)
	return key, nil
}

// Generate matching key pairs for Alice and Bob, returning a symmetric key as well
func DHKeyPairs(p, g *big.Int) ([2]*big.Int, [2]*big.Int, [32]byte, error) {
	// Alice's key pair
	a, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		return [2]*big.Int{}, [2]*big.Int{}, [32]byte{}, err
	}
	a.Add(a, big.NewInt(1))
	A := ModExp(g, a, p)
	aliceKeyPair := [2]*big.Int{a, A}

	// Bob's key pair
	b, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		return [2]*big.Int{}, [2]*big.Int{}, [32]byte{}, err
	}
	b.Add(b, big.NewInt(1))
	B := ModExp(g, b, p)
	bobKeyPair := [2]*big.Int{b, B}

	//Derive Symmetric key
	symmetricKey, err := DHDeriveSymmetricKeys(aliceKeyPair, bobKeyPair, p)
	if err != nil {
		return [2]*big.Int{}, [2]*big.Int{}, [32]byte{}, err
	}
	return aliceKeyPair, bobKeyPair, symmetricKey, nil
}

// Generate one key pair
func DHKeyPair(p, g *big.Int) (*big.Int, *big.Int, error) {
	priv, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		return nil, nil, err
	}
	priv.Add(priv, big.NewInt(1))
	pub := ModExp(g, priv, p)

	return priv, pub, nil
}
