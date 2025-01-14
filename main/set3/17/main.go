package main

// CBC padding oracle
import (
	"crypto-pals/crysuite"
	"crypto-pals/util"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

var key []byte
var iv []byte

func choose_str() ([]byte, string) {
	data, _ := os.ReadFile("data.txt")
	strArray := strings.Split(string(data), "\n")

	rand.Seed(uint64(time.Now().UnixNano()))
	rndIndex := rand.Intn(len(strArray))
	rndStr := util.DecodeB64(strArray[rndIndex])

	enc, _ := crysuite.EncryptAES_CBC(rndStr, key, iv)

	return enc, string(rndStr)
}

func valid_pad(enc []byte) bool {
	dec, _ := crysuite.DecryptAES_CBC(enc, key, iv)
	veri := util.DetectPad(string(dec), len(key))

	return veri
}

// Format encrypted data into keysize blocks
func breakOracle(enc []byte, iv []byte) ([]byte, error) {
	var pt []byte
	if !valid_pad(enc) {
		panic("Text has invalid padding")
	}

	chunkEnc := util.Chunkify(enc, len(key))
	lastBlock := iv

	for _, chunk := range chunkEnc {
		ptBlock, err := break_block(chunk, lastBlock)
		if err != nil {
			return []byte(""), err
		}
		pt = append(pt, ptBlock...)
		lastBlock = chunk
	}

	return pt, nil
}

// Iterating through each block
func break_block(block, lastBlock []byte) ([]byte, error) {
	var pt []byte

	for n := range block {
		ptByte, err := break_char(block, lastBlock, pt, n+1)
		if err != nil {
			return []byte(""), err
		}
		pt = append([]byte{ptByte}, pt...)
	}

	return pt, nil
}

// Guess each byte, check response
func break_char(block, lastBlock, knownPt []byte, byteN int) (byte, error) {
	var blockSuff []byte

	offset := len(lastBlock) - byteN
	encByte := lastBlock[offset]
	blockPref := make([]byte, offset)

	copy(blockPref, lastBlock[:offset])

	for i, v := range knownPt {
		blockSuff = append(blockSuff, v^byte(byteN)^lastBlock[(len(lastBlock))-(len(knownPt)-i)])
	}

	for b := byte(0); ; b++ {
		nBlock := append(append(append(blockPref, b), blockSuff...), block...)
		if valid_pad(nBlock) {
			if b == encByte && byteN == 1 {
				continue
			}

			return b ^ encByte ^ byte(byteN), nil
		}
		if b == 255 {
			break
		}
	}

	return 0, errors.New("no valid byte found")
}

func main() {
	key = util.GenerateRandomBytes(16)
	iv = util.GenerateRandomBytes(16)

	enc, str := choose_str()
	verify := valid_pad(enc)

	fmt.Println(verify)

	dec, _ := breakOracle(enc, iv)
	fmt.Println(str)
	fmt.Println(string(dec))
}
