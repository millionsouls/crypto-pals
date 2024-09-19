package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type Bytes []byte

func DecodeHex(hexstr string) []byte {
	bytes, err := hex.DecodeString(hexstr)

	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

func Encodeb64(bytes []byte) string {
	b64 := base64.StdEncoding.EncodeToString(bytes)
	return b64
}

func Xor(key byte, b Bytes) Bytes {
	result := make(Bytes, len(b))
	for i := range b {
		result[i] = b[i] ^ key
	}
	return result
}
