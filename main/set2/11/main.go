package main

//ECB/CBC detection oracle
import (
	"crypto-pals/lib/crysuite"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return b
}

func encryptionOracle(input []byte) ([]byte, []byte) {
	var out []byte
	choice, err := rand.Int(rand.Reader, big.NewInt(2))
	key := GenerateRandomBytes(16)

	if err != nil {
		panic(err)
	}

	switch choice.Int64() {
	case 0:
		fmt.Println("Using ECB")
		out, _ = crysuite.EncryptAES_ECB(input, key)
	case 1:
		fmt.Println("Using CBC")
		out, _ = crysuite.EncryptAES_CBC(input, key, []byte("\x00"))
	}

	return out, key
}

func main() {
	//text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	text, err := os.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	encrypted, key := encryptionOracle(text)
	test := crysuite.DetectECB(encrypted, len(key))

	if test {
		fmt.Println("Is ECB")
		decrypted, err := crysuite.DecryptAES_ECB(encrypted, key)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(decrypted))
	} else {
		fmt.Println("Is CBC")
		decrypted, err := crysuite.DecryptAES_CBC(encrypted, key, []byte(""))
		if err != nil {
			panic(err)
		}
		fmt.Println(string(decrypted))
	}
}
