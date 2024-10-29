package main

// CBC bitflipping attack
import (
	"crypto-pals/util"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var setKey sync.Once
var key []byte
var iv []byte

func wrap_data(input string) []byte {
	prefix := []byte("comment1=cooking%20MCs;userdata=")
	suffix := []byte(";comment2=%20like%20a%20pound%20of%20bacon")

	re := regexp.MustCompile(`[;=]`)
	text := []byte(re.ReplaceAllString(input, ""))
	text = append(prefix, text...)
	text = append(text, suffix...)

	return util.AESCBCEncrypt(text, key, iv)
}

func isAdmin(input []byte) bool {
	dec := util.AESCBCDecrypt(input, key, iv)
	splits := strings.Split(string(dec), ";")
	result := make(map[string]string)

	for _, j := range splits {
		kv := strings.Split(j, "=")
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		} else {
			fmt.Println(kv)
			panic("Key value unexpected length")
		}
	}

	return result["admin"] == "true"
}

func main() {
	setKey.Do(func() { key = util.GenerateRandomBytes(16); iv = util.GenerateRandomBytes(16) })

	input := ";admin=true;"
	enc := wrap_data(input)

	isAd := isAdmin(enc)
	fmt.Println(isAd)
}
