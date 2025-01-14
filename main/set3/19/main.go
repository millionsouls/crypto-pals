package main

import (
	"bytes"
	"crypto-pals/crysuite"
	"crypto-pals/util"
	"fmt"
	"os"
	"sort"
	"strings"
)

var key []byte

func enc(strArray []string) [][]byte {
	bArrArr := make([][]byte, len(strArray))

	for i, val := range strArray {
		enc := crysuite.AES_CTR_Encrypt(util.DecodeB64(val), key, bytes.Repeat([]byte("\x00"), 8))
		bArrArr[i] = enc
	}

	sort.Slice(bArrArr, func(i, j int) bool { return len(bArrArr[i]) < len(bArrArr[j]) })

	return bArrArr
}

func dec(enc [][]byte) ([]string, []byte) {
	var trunc []byte
	min := len(enc[0])

	for _, v := range enc {
		trunc = append(trunc, v[:min]...)
	}

	dec, keyStream := util.ComputeKey(trunc, util.FindKeySize(trunc))
	split := make([]string, len(enc))

	start := 0
	end := min

	for i := 0; i < len(enc); i++ {
		split[i] = string(dec[start:end])
		start = end
		end = start + min
	}

	return split, keyStream
}

func main() {
	key = util.GenerateRandomBytes(16)
	// nonce := 0

	data, _ := os.ReadFile("data.txt")
	strArray := strings.Split(string(data), "\n")

	enc := enc(strArray)
	dec, keys := dec(enc)

	fmt.Println(string(keys))
	fmt.Println(strings.Join(dec, "\n"))
}
