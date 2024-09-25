package main

import (
	"crypto-pals/util"
	"encoding/hex"
	"fmt"
)

func main() {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"

	var d1 []byte = util.DecodeHex(hex1)
	var d2 []byte = util.DecodeHex(hex2)

	if len(hex1) != len(hex2) {
		fmt.Println("Buffers must have equal length")
	}

	result := make([]byte, len(d1))
	for i := 0; i < len(d1); i++ {
		result[i] = d1[i] ^ d2[i]
	}
	hexresult := hex.EncodeToString(result)

	fmt.Println(hexresult)
}
