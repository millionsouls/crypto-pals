package main

import (
	"bytes"
	"crypto-pals/util"
	"fmt"
	"os"
)

var key []byte
var mText []byte

func theOracle(input []byte) []byte {
	return util.AESEncrypt(input, key)
}

func findKeySize() int {
	pLen := 0

	for size := 1; size <= 512; size++ { // Arbitrarily chosen max size
		input := bytes.Repeat([]byte("A"), size)
		encrypted := theOracle(input)

		// Check the length of the encrypted message
		if len(encrypted) > pLen {
			if pLen != 0 {
				return pLen
			}
			pLen = len(encrypted)
		}
	}
	return 0
}

func genByteTable(prefix []byte, start, end int) map[string]byte {
	var i byte
	result := make(map[string]byte)

	for i = 1; i < 255; i++ {
		msg := append(prefix, i)
		test := theOracle(append(msg, mText...))[start:end]

		result[string(test)] = i
	}

	return result
}

func findKey(key int) {
	emptyCipher := theOracle(mText)
	var decrypted []byte

	for len(decrypted) < len(emptyCipher) {
		blockStart := len(decrypted)
		blockEnd := blockStart + key

		for i := key - 1; i >= 0; i-- {
			prefix := bytes.Repeat([]byte("A"), i)
			known := append(prefix, decrypted...)
			table := genByteTable(known, blockStart, blockEnd)

			block := theOracle(append(prefix, mText...))[blockStart:blockEnd]
			decrypted = append(decrypted, table[string(block)])
		}
	}

	fmt.Printf("Decrypted string: %s\n", decrypted)
}

func main() {
	text, _ := os.ReadFile("data.txt")
	mmText, _ := os.ReadFile("../11/data.txt") // text from challenge 11

	mText = util.DecodeB64(string(text))
	key = util.GenerateRandomBytes(16)

	fmt.Println("Setting globals")

	encrypted := theOracle(mmText)
	size := findKeySize()

	fmt.Println("Encrypting test message and finding the key length")
	fmt.Println(size)

	isECB := util.DetectECB(encrypted, size)
	if !isECB {
		panic("Detected message was not ECB")
	}

	fmt.Println("Finding the key")

	findKey(size)
}
