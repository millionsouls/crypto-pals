package main

// Repeating-key XOR
import (
	"crypto-pals/util"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	newStr, _ := util.Xor([]byte("ICE"), []byte(str))

	fStr := hex.EncodeToString(newStr)

	fmt.Println(fStr)
}
