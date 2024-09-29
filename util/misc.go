package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"unicode"
)

// genertic
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

func DecodeB64(str string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return bytes
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

func HamDis(one []byte, two []byte) int {
	oneLen := len(one)
	twoLen := len(two)
	diff := 0

	if oneLen != twoLen {
		panic("Hamming: Inputs of different lengths")
	}

	for i, j := range one {
		nb := two[i]
		for k := 1; k < 129; k = 2 * k {
			if (j & byte(k)) != (nb & byte(k)) {
				diff++
			}
		}
	}

	return diff
}
