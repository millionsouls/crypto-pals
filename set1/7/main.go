package main

import (
	"crypto-pals/util"
	"crypto/aes"
	"fmt"
	"os"
)

func UnPKCS7(pt []byte) []byte {
	return pt[:len(pt)-int(pt[len(pt)-1])]
}

func aesDecrypt(data, key []byte) []byte {
	cipher, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	decrypted := make([]byte, len(data))
	size := cipher.BlockSize()
	chunks := util.Chunkify(data, size)

	for i, chunk := range chunks {
		cipher.Decrypt(decrypted[i*size:(i+1)*size], chunk)
	}

	return UnPKCS7(decrypted)
}

func main() {
	data, err := os.ReadFile("data.txt")
	key := []byte("YELLOW SUBMARINE")

	if err != nil {
		panic(err)
	}

	ddata := util.DecodeB64(string(data))
	text := aesDecrypt(ddata, key)

	fmt.Println(string(text))
}
