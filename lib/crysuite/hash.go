package crysuite

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func HMAC256(key, data []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func VerifyHMAC(mac1, mac2 []byte) bool {
	return hmac.Equal(mac1, mac2)
}
