package main

// CBC Mode
import (
	"crypto-pals/util"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	key := []byte("YELLOW SUBMARINE")
	iv := []byte("\x00")

	if err != nil {
		panic(err)
	}

	text := util.AES_CBC_Decrypt(util.DecodeB64(string(data)), key, iv)

	fmt.Println(string(text))
}
