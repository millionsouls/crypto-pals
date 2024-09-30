package util

import (
	"crypto/aes"
	"fmt"
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

func AESEncrypt(data, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	size := cipher.BlockSize()
	chunks := Chunkify(data, size)
	chunks = append(chunks[:len(chunks)-1], PKCS7(chunks[len(chunks)-1], len(key)))
	encrypted := make([]byte, len(chunks)*size)

	for i, chunk := range chunks {
		cipher.Encrypt(encrypted[i*size:(i+1)*size], chunk)
	}

	fmt.Println(encrypted)
	fmt.Println(string(AESDecrypt(encrypted, key)))

	return encrypted
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

func AESCBCEncrypt(data, key, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	size := cipher.BlockSize()

	data = PKCS7(data, size)
	chunks := Chunkify(data, size)

	//chunks := Chunkify(data, size)
	//chunks = append(chunks[:len(chunks)-1], PKCS7(chunks[len(chunks)-1], len(key)))

	encrypted := make([]byte, len(chunks)*size)

	lastChunk := iv
	for i, chunk := range chunks {
		cipher.Encrypt(encrypted[i*size:(i+1)*size], RXor(chunk, lastChunk))
		lastChunk = encrypted[i*size : (i+1)*size]
	}

	fmt.Println(encrypted)
	fmt.Println(AESCBCDecrypt(encrypted, key, []byte("\x00")))
	return encrypted
}
