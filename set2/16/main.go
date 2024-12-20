package main

// CBC bitflipping attack
import (
	"bytes"
	"crypto-pals/util"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var setKey sync.Once
var key []byte
var iv []byte

func wrap_data(input []byte) []byte {
	prefix := []byte("comment1=cooking%20MCs;userdata=")
	suffix := []byte(";comment2=%20like%20a%20pound%20of%20bacon")

	re := regexp.MustCompile(`[;=]`)
	text := []byte(re.ReplaceAllString(string(input), ""))
	text = append(prefix, text...)
	text = append(text, suffix...)

	return util.AES_CBC_Encrypt(text, key, iv)
}

func isAdmin(input []byte) bool {
	dec := util.AES_CBC_Decrypt(input, key, iv)
	splits := strings.Split(string(dec), ";")
	result := make(map[string]string)

	fmt.Println(string(dec))

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

func make_admin() []byte {
	block := bytes.Repeat([]byte("A"), 16)
	ct := wrap_data(append(block, block...))

	// fmt.Println(isAdmin(ct))

	flip := util.RXor(block, append([]byte(";admin=true"), bytes.Repeat([]byte("A"), 11-len(";admin=true"))...))
	pad := append(bytes.Repeat([]byte("\x00"), 16*3-len(flip)), flip...)
	pad = append(pad, bytes.Repeat([]byte("\x00"), len(ct)-len(pad))...)
	new_ct := util.RXor(ct, pad)

	return new_ct
}

func main() {
	setKey.Do(func() { key = util.GenerateRandomBytes(16); iv = util.GenerateRandomBytes(16) })

	input := ";admin=true;"
	enc := wrap_data([]byte(input))

	fmt.Println(isAdmin(enc))

	test := make_admin()
	fmt.Println(isAdmin(test))
}
