package util

// functions for AES cipher suite; ECB, CBC
// detectecb - finds any repeat cipher blocks which usually comes from ECB methods

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
)

func DetectECB(data []byte, size int) bool {
	chunks := Chunkify(data, size)
	// chunks = append(chunks[:len(chunks)-1], PKCS7(chunks[len(chunks)-1], size))

	chunkFreq := make(map[string]int)
	// repeats := 0

	for _, chunk := range chunks {
		chunkStr := string(chunk)
		chunkFreq[chunkStr]++
		if chunkFreq[chunkStr] > 1 {
			return true
		}
	}

	return false
}

func AES_ECB_Decrypt(data, key []byte) []byte {
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

func AES_ECB_Encrypt(data, key []byte) []byte {
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

	//fmt.Println(encrypted)
	//fmt.Println(string(AESDecrypt(encrypted, key)))

	return encrypted
}

func AES_CBC_Decrypt(data []byte, key []byte, iv []byte) []byte {
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

func AES_CBC_Encrypt(data, key, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	size := cipher.BlockSize()
	//data = PKCS7(data, len(key))
	chunks := Chunkify(data, size)
	chunks = append(chunks[:len(chunks)-1], PKCS7(chunks[len(chunks)-1], len(key)))

	encrypted := make([]byte, len(chunks)*size)

	lastChunk := iv
	for i, chunk := range chunks {
		cipher.Encrypt(encrypted[i*size:(i+1)*size], RXor(lastChunk, chunk))
		lastChunk = encrypted[i*size : (i+1)*size]
	}

	//fmt.Println(encrypted)
	//fmt.Println(string(AESCBCDecrypt(encrypted, key, []byte("\x00"))))

	return encrypted
}

func AES_CTR_Decrypt(pt, key []byte, nonce uint64) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	nonceBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceBytes, nonce)

	plaintext := make([]byte, len(pt))

	for i := 0; i < len(pt); i += aes.BlockSize {
		//create a new iv for each block
		iv := append(nonceBytes, make([]byte, 8)...)                   // 8 bytes for the nonce and 8 for the counter
		binary.LittleEndian.PutUint64(iv[8:], uint64(i/aes.BlockSize)) // set the counter

		//initialize ctr for this block only
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(plaintext[i:], pt[i:])
	}

	return plaintext
}
