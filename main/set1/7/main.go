package main

// AES ECB decryption
import (
	"crypto-pals/util"
	"crypto/aes"
	"fmt"
	"os"
)

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

	return decrypted
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
