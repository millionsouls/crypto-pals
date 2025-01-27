package main

// CBC Mode
import (
	"crypto-pals/lib/crysuite"
	"crypto-pals/lib/util"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	key := []byte("YELLOW SUBMARINE")

	if err != nil {
		panic(err)
	}

	text, _ := crysuite.DecryptAES_CBC(util.DecodeB64(string(data)), key, []byte(""))

	fmt.Println(string(text))
}
