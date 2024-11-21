package main

import (
	"crypto-pals/util"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

var key []byte

func dec_ctr(pt, key []byte, nonce uint64) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	nonceBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceBytes, nonce)
	plaintext := make([]byte, len(pt))

	for i := 0; i < len(pt); i += aes.BlockSize {
		iv := append(make([]byte, 8), nonceBytes...)
		binary.LittleEndian.PutUint64(iv[8:], uint64(i/aes.BlockSize)) // set the counter

		//initialize ctr for this block only
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(plaintext[i:], pt[i:])
	}

	return plaintext
}

func enc_ctr(pt, key []byte, nonce uint64) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	nonceBytes := make([]byte, 16)
	binary.LittleEndian.PutUint64(nonceBytes, nonce)
	plaintext := make([]byte, len(pt))

	stream := cipher.NewCTR(block, nonceBytes)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return append(nonceBytes, ciphertext...)
}

func main() {
	key = util.GenerateRandomBytes(16)
	nonce := 0

	data, _ := os.ReadFile("data.txt")
	strArray := strings.Split(string(data), "\n")

	for _, i := range strArray {
		// fmt.Println(string(util.DecodeB64(i)))
		enc := enc_ctr(util.DecodeB64(i), key, uint64(nonce))
		dec := dec_ctr(enc, key, uint64(nonce))

		fmt.Println(string(dec))
	}

}
