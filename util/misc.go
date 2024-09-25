package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"unicode"
)

func DecodeHex(hexstr string) []byte {
	bytes, err := hex.DecodeString(hexstr)

	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

func EncodeB64(bytes []byte) string {
	b64 := base64.StdEncoding.EncodeToString(bytes)
	return b64
}

func CleanText(text []byte) []byte {
	var cleaned []byte
	for _, b := range text {
		if unicode.IsPrint(rune(b)) {
			cleaned = append(cleaned, b)
		}
	}
	return cleaned
}
