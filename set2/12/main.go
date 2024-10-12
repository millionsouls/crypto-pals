package main

import (
	"bytes"
	"crypto-pals/util"
	"os"
)

var key []byte

func theOracle(input []byte) []byte {
	return util.AESEncrypt(input, key)
}

func findBlockSize() int {
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

func findKey() {

}

func main() {
	text, _ := os.ReadFile("data.txt")
	mText, _ := os.ReadFile("../11/data.txt") // text from challenge 11

	dText := util.DecodeB64(string(text))
	mText = append(mText, dText...)

	key = util.GenerateRandomBytes(16)

	encrypted := theOracle(mText)
	size := findBlockSize()

	isECB := util.DetectECB(encrypted, size)

	if !isECB {
		panic("Detected message was not ECB")
	}
}
