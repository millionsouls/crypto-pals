package main

// AES-CTR
import (
	"crypto-pals/util"
	"fmt"
)

func main() {
	text := "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
	key := "YELLOW SUBMARINE"

	nonce := uint64(0)
	enc := util.DecodeB64(text)
	dec := util.AES_CTR_Decrypt(enc, []byte(key), nonce)

	fmt.Println(string(dec))
}
