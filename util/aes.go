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
	var decrypted []byte
	lastchunk := iv

	for _, chunk := range chunks {
		decChunk := make([]byte, size)
		cipher.Decrypt(decChunk, chunk)
		decrypted = append(decrypted, RXor(lastchunk, decChunk)...)
		lastchunk = chunk
	}

	return decrypted
}
