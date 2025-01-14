package main

// AES-CTR
import (
	"crypto-pals/crysuite"
	"crypto-pals/util"
	"fmt"
)

func main() {
	text := "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
	key := "YELLOW SUBMARINE"

	nonce := uint64(0)
	enc := util.DecodeB64(text)
	dec, _ := crysuite.DecryptAES_CTR(enc, []byte(key), nonce)

	fmt.Println(string(dec))
}
