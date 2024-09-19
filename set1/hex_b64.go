package set1

import (
	"encoding/hex"
	"fmt"

	"crypto-tools/util"
)

func Hexb64() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(util.Encodeb64(bytes))
}
