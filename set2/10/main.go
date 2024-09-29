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

	text := util.AESCBCDecrypt(data, key, iv)

	fmt.Println(text)
}
