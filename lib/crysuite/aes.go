package crysuite

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"strings"

	"crypto-pals/lib/util"
)

// DetectECB checks if a ciphertext uses ECB mode by finding repeating blocks
func DetectECB(data []byte, blockSize int) bool {
	chunks := util.Chunkify(data, blockSize)

	chunkFreq := make(map[string]int)
	for _, chunk := range chunks {
		chunkStr := string(chunk)
		chunkFreq[chunkStr]++
		if chunkFreq[chunkStr] > 1 {
			return true
		}
	}

	return false
}

// DecryptAES_ECB decrypts the given data using AES in ECB mode
func DecryptAES_ECB(data, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := cipher.BlockSize()
	chunks := util.Chunkify(data, size)

	decrypted := make([]byte, len(data))
	for i, chunk := range chunks {
		cipher.Decrypt(decrypted[i*size:(i+1)*size], chunk)
	}

	return decrypted, nil
}

// EncryptAES_ECB encrypts the given data using AES in ECB mode
func EncryptAES_ECB(data, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := cipher.BlockSize()
	chunks := util.Chunkify(data, size)

	// Apply PKCS7 padding to the last chunk
	if len(chunks) > 0 {
		chunks[len(chunks)-1] = util.PKCS7(chunks[len(chunks)-1], len(key))
	}

	encrypted := make([]byte, len(data))
	for i, chunk := range chunks {
		cipher.Encrypt(encrypted[i*size:(i+1)*size], chunk)
	}

	return encrypted, nil
}

// DecryptAES_CBC decrypts the given data using AES in CBC mode
func DecryptAES_CBC(data, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	var decrypted []byte
	size := cipher.BlockSize()
	chunks := util.Chunkify(data[size:], size)
	iiv := iv
	if len(iv) == 0 {
		iiv = data[:size]
	}

	for _, chunk := range chunks {
		decChunk := make([]byte, size)
		cipher.Decrypt(decChunk, chunk)
		xoredChunk, err := util.Xor(iiv, decChunk)
		if err != nil {
			return nil, err
		}
		decrypted = append(decrypted, xoredChunk...)
		iiv = chunk
	}

	// Remove padding
	padding := int(decrypted[len(decrypted)-1])
	return decrypted[:len(decrypted)-padding], nil
}

// EncryptAES_CBC encrypts the given data using AES in CBC mode
func EncryptAES_CBC(data, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := cipher.BlockSize()
	chunks := util.Chunkify(data, size)
	chunks[len(chunks)-1] = util.PKCS7(chunks[len(chunks)-1], size)
	encrypted := make([]byte, len(chunks)*size)
	lastChunk := iv

	if len(iv) == 0 {
		iv = make([]byte, size)
		lastChunk = iv
	}

	// Loop through each chunk and encrypt it
	for i, chunk := range chunks {
		for j := 0; j < size; j++ {
			chunk[j] ^= lastChunk[j]
		}

		// Encrypt the XORed chunk
		cipher.Encrypt(encrypted[i*size:(i+1)*size], chunk)

		// Update lastChunk for the next block (current ciphertext)
		lastChunk = encrypted[i*size : (i+1)*size]
	}

	return append(iv, encrypted...), nil
}

// DecryptAES_CTR decrypts the given data using AES in CTR mode
func DecryptAES_CTR(data, key []byte, nonce uint64) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonceBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceBytes, nonce)

	plaintext := make([]byte, len(data))
	for i := 0; i < len(data); i += aes.BlockSize {
		// Create a new IV for this block: nonce + counter
		iv := append(nonceBytes, make([]byte, 8)...)                   // 8 bytes for nonce and counter
		binary.LittleEndian.PutUint64(iv[8:], uint64(i/aes.BlockSize)) // Set the counter

		// Initialize CTR mode stream cipher with the new IV
		stream := cipher.NewCTR(cipherBlock, iv)
		stream.XORKeyStream(plaintext[i:], data[i:])
	}

	return plaintext, nil
}

// EncryptAES_CTR encrypts the given data using AES in CTR mode
func EncryptAES_CTR(data, key, nonce []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := cipherBlock.BlockSize()
	if len(nonce) != blockSize/2 {
		return nil, errors.New("invalid nonce size")
	}

	chunks := util.Chunkify(data, blockSize)
	encryptedChunks := make([]string, len(chunks))

	// Process chunks sequentially
	for i, chunk := range chunks {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(i))
		iv := append(nonce, b...)
		ks, err := EncryptAES_ECB(iv, key) // Reuse the ECB encryption for CTR

		if err != nil {
			return nil, err
		}

		// XOR the chunk with the keystream
		xoredChunk, err := util.Xor(chunk, ks)
		if err != nil {
			return nil, err
		}
		encryptedChunks[i] = string(xoredChunk)
	}

	return []byte(strings.Join(encryptedChunks, "")), nil
}
