package main

import (
	"bytes"
	"crypto-pals/crysuite"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func main() {
	// Create base and modulus
	p, _ := rand.Int(rand.Reader, new(big.Int).Sub(big.NewInt(100), big.NewInt(1)))
	g := big.NewInt(5)
	// Alice and Bob generates key pair
	a, _, _ := crysuite.DHKeyPair(p, g)
	b, _, _ := crysuite.DHKeyPair(p, g)
	m, _, _ := crysuite.DHKeyPair(p, g)
	// MITM -> Bob / p g p
	// MITM -> Alice / p g p
	sAlice, _, _ := crysuite.DHDeriveSymmetricKey(p, a, p)
	sBob, _, _ := crysuite.DHDeriveSymmetricKey(p, b, p)
	sMITM, _, _ := crysuite.DHDeriveSymmetricKey(p, m, p)

	fmt.Printf("Alice: %x\nBob: %x\nMITM: %x\n", sAlice, sBob, sMITM)

	// Test keys
	if !bytes.Equal(sAlice[:], sBob[:]) || !bytes.Equal(sMITM[:], sBob[:]) {
		fmt.Println("***Shared secret mismatch***")
		return
	}
	fmt.Println("***Shared secret match***")

	// Message tests
	plaintext, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	ciphertext, err := crysuite.EncryptAES_CBC(plaintext, sAlice[:], []byte{})
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}
	decryptedMessage, err := crysuite.DecryptAES_CBC(ciphertext, sMITM[:], []byte{})
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Printf("Decrypted: %s\n", string(decryptedMessage))
}
