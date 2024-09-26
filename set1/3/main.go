package main

import (
	"fmt"
	"math"
	"unicode"

	"crypto-pals/util"
)

var freqTable = map[rune]float64{
	'a': 0.0817, 'b': 0.0149, 'c': 0.0278, 'd': 0.0425, 'e': 0.1270,
	'f': 0.0223, 'g': 0.0202, 'h': 0.0609, 'i': 0.0697, 'j': 0.0015,
	'k': 0.0077, 'l': 0.0403, 'm': 0.0241, 'n': 0.0675, 'o': 0.0751,
	'p': 0.0193, 'q': 0.0010, 'r': 0.0599, 's': 0.0633, 't': 0.0906,
	'u': 0.0276, 'v': 0.0098, 'w': 0.0236, 'x': 0.0015, 'y': 0.0197,
	'z': 0.0007,
}

func chiSquaredScore(text []byte) float64 { // to util
	frequency := make(map[rune]int)
	totalLetters := 0

	// count letter frequencies
	for _, b := range text {
		if unicode.IsLetter(rune(b)) {
			lower := unicode.ToLower(rune(b))
			frequency[lower]++
			totalLetters++
		}
	}

	// calculate the Chi-squared score
	score := 0.0
	if totalLetters == 0 {
		return math.MaxFloat64 // penalize empty results
	}

	for letter, count := range frequency {
		expected := freqTable[letter] * float64(totalLetters)
		if expected > 0 {
			score += math.Pow(float64(count)-expected, 2) / expected
		}
	}

	return score
}

func cleanText(text []byte) []byte { // to util
	var cleaned []byte
	for _, b := range text {
		if unicode.IsPrint(rune(b)) {
			cleaned = append(cleaned, b)
		}
	}
	return cleaned
}

func main() {
	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// hex to bytes
	data := util.DecodeHex(hexStr)
	bestScore := math.MaxFloat64
	bestChar := byte(0)
	bestResult := []byte{}

	// test single-byte characters
	for i := 0; i <= 255; i++ {
		xored := util.Xor(byte(i), data)
		cleaned := cleanText(xored)
		score := chiSquaredScore(cleaned)

		if score < bestScore && score != 0 {
			bestScore = score
			bestChar = byte(i)
			bestResult = cleaned
		}
	}

	fmt.Printf("Best XOR character: '%c' (0x%X)\n", bestChar, bestChar)
	fmt.Printf("XORed result: '%s'\n", string(bestResult))
	fmt.Printf("Best Score: %f\n", bestScore)

}
