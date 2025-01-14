package main

import (
	"crypto-pals/crysuite"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
)

func modExp(base, exp, mod *big.Int) *big.Int {
	result := new(big.Int).SetInt64(1)
	baseMod := new(big.Int).Mod(base, mod)

	expBinary := fmt.Sprintf("%b", exp)

	for i := 0; i < len(expBinary); i++ {
		result = new(big.Int).Mod(new(big.Int).Mul(result, result), mod)
		if expBinary[i] == '1' {
			result = new(big.Int).Mod(new(big.Int).Mul(result, baseMod), mod)
		}
	}
	return result
}

func generatePrivateKey(p *big.Int) (*big.Int, error) {
	privKey, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1))) // between 0 and p-1
	if err != nil {
		return nil, err
	}
	privKey.Add(privKey, big.NewInt(1)) // make sure private key is between 1 and p-1
	return privKey, nil
}

func bigIntToBinary(n *big.Int) string {
	return fmt.Sprintf("%b", n)
}

func main() {
	//Generate key pairs
	p, _ := new(big.Int).SetString("ffffffffffffffffc90fdaa22168c234c4c6628b80dc1cd129024e088a67cc74020bbea63b139b22514a08798e3404ddef9519b3cd3a431b302b0a6df25f14374fe1356d6d51c245e485b576625e7ec6f44c42e9a637ed6b0bff5cb6f406b7edee386bfb5a899fa5ae9f24117c4b1fe649286651ece45b3dc2007cb8a163bf0598da48361c55d39a69163fa8fd24cf5f83655d23dca3ad961c62f356208552bb9ed529077096966d670c354e4abc9804f1746c08ca237327fff", 16)
	g := big.NewInt(2) // g = 2

	a, err := generatePrivateKey(p)
	if err != nil {
		fmt.Println("Error generating private key for Alice:", err)
		return
	}
	b, err := generatePrivateKey(p)
	if err != nil {
		fmt.Println("Error generating private key for Bob:", err)
		return
	}

	fmt.Printf("Alice private: %s\n", a.Text(16))
	fmt.Printf("Bob private: %s\n", b.Text(16))

	A := modExp(g, a, p)
	B := modExp(g, b, p)

	fmt.Printf("Alice public: %s\n", A.Text(16))
	fmt.Printf("Bob public: %s\n", B.Text(16))

	sAlice := modExp(B, a, p)
	sBob := modExp(A, b, p)

	if sAlice.Cmp(sBob) != 0 {
		fmt.Println("***Shared secrets do not match!***")
		return
	} else {
		fmt.Println("***Shared secrets match!***")
	}

	sharedSecret := sAlice.Bytes()
	key := sha256.Sum256(sharedSecret)

	fmt.Printf("SS Alice: %s\n", sAlice.Text(16))
	fmt.Printf("SS Alice: %s\n", sBob.Text(16))
	fmt.Printf("***Symmetric key***: %x\n", key)

	// Test message
	plaintext, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	ciphertext, err := crysuite.EncryptAES_CBC(plaintext, key[:], []byte{})
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}

	fmt.Printf("Encrypted: %x\n", ciphertext)
	decryptedMessage, err := crysuite.DecryptAES_CBC(ciphertext, key[:], []byte{})

	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Printf("Decrypted: %s\n", string(decryptedMessage))
}
