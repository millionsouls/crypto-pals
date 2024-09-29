package util

import (
	"crypto/aes"
)

func AESDecrypt(data, key []byte) []byte {
	cipher, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	decrypted := make([]byte, len(data))
	size := cipher.BlockSize()
	chunks := Chunkify(data, size)

	for i, chunk := range chunks {
		cipher.Decrypt(decrypted[i*size:(i+1)*size], chunk)
	}

	return decrypted
}

func AESCBCDecrypt(data []byte, key []byte, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	size := cipher.BlockSize()
	chunks := Chunkify(data, size)
	decrypted := make([]byte, len(data))
	lastchunk := iv

	for _, chunk := range chunks {

		lastchunk = chunk
	}

	return decrypted
}
