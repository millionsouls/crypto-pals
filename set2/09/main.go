package main

// PKCS#7 Padding

import (
	"crypto-pals/util"
	"fmt"
)

func main() {
	text := "YELLOW SUBMARINE"

	ntext := util.PKCS7([]byte(text), 20)
	untext := util.UnPad(ntext)

	fmt.Println(ntext)
	fmt.Println(untext)
}
