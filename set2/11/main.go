package main

import (
	"crypto-pals/util"
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return b
}

func encryptionOracle(input []byte) []byte {
	var out []byte
	choice, err := rand.Int(rand.Reader, big.NewInt(2))
	key := GenerateRandomBytes(16)

	if err != nil {
		panic(err)
	}

	switch choice.Int64() {
	case 0:
		fmt.Println("Using ECB")
		out = util.AESEncrypt(input, key)
	case 1:
		fmt.Println("Using CBC")
		out = util.AESCBCEncrypt(input, key, []byte("\x00"))
	}

	return out
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	encryptionOracle([]byte(text))
}
