package set1

import (
	"crypto-tools/util"
	"fmt"
	"math"
)

type Bytes []byte

func XorCipherBrute(encode string) {
	data := util.DecodeHex(encode)

	for i := 'a'; i <= 'z'; i++ {
		result := make([]byte, len(data))

		for j := range data {
			result[j] = data[j] ^ byte(i)
		}

		resultStr := string(result)

		fmt.Printf("%c: %s\n", i, resultStr)
	}
}

func scoreFrequencyAnalysis(data []byte) float64 {
	// English letter frequencies (you can refine this)
	frequency := map[byte]float64{
		'e': 0.12702, 't': 0.09056, 'a': 0.08167, 'o': 0.07507,
		'i': 0.06966, 'n': 0.06749, 's': 0.06327, 'h': 0.06155,
		'r': 0.05987, 'd': 0.04253, 'l': 0.04025, 'c': 0.02782,
		'u': 0.02758, 'm': 0.02406, 'w': 0.02360, 'f': 0.02228,
		'g': 0.02015, 'y': 0.01974, 'p': 0.01929, 'b': 0.01492,
		'v': 0.00978, 'k': 0.00849, 'j': 0.00153, 'x': 0.00150,
		'q': 0.00095, 'z': 0.00074,
	}

	score := 0.0
	for _, b := range data {
		if val, exists := frequency[b]; exists {
			score += val
		}
	}
	return score
}

func (b Bytes) xor(key byte) Bytes {
	result := make(Bytes, len(b))
	for i := range b {
		result[i] = b[i] ^ key
	}
	return result
}

func AttemptSingleByteXor(cipher Bytes) (byte, Bytes) {
	bestKey := byte(0)
	bestScore := math.Inf(-1)
	var decrypted Bytes

	for key := 0; key < 256; key++ {
		plaintext := cipher.xor(byte(key))
		score := scoreFrequencyAnalysis(plaintext)

		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			decrypted = plaintext
		}
	}

	return bestKey, decrypted
}
