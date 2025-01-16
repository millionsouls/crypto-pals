package main

// Byte at a time ECB decryption
import (
	"bytes"
	"crypto-pals/lib/crysuite"
	"crypto-pals/lib/util"
	"fmt"
	"os"

	"golang.org/x/exp/rand"
)

var key []byte

func theOracle(input []byte) []byte {
	encrypted, err := crysuite.EncryptAES_ECB(input, key)
	if err != nil {
		panic(err)
	}
	return encrypted
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

func randomOracle(atkText []byte) []byte {
	text, _ := os.ReadFile("../12/data.txt") // text from challenge 12
	decoded := util.DecodeB64(string(text))

	numBytes := rand.Intn(256)
	randPref := util.GenerateRandomBytes(numBytes)

	atkText = append(randPref, atkText...)
	atkText = append(atkText, decoded...)

	return atkText
}

func main() {
	fmt.Println("Globals setup")
	key = util.GenerateRandomBytes(16)
	encrypted := randomOracle(util.GenerateRandomBytes(128))

	fmt.Println("Finding key size")
	findKeySize()

	isECB := crysuite.DetectECB(encrypted, len(key))
	fmt.Println(isECB)

}
