package main

import (
	"crypto-pals/util"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	newStr := util.RXor([]byte("ICE"), []byte(str))

	fStr := hex.EncodeToString(newStr)

	fmt.Println(fStr)
}
