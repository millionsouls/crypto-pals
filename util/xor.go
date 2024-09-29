package util

import (
	"math"
	"unicode"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

var FreqTable = map[rune]float64{
	'a': 0.0817, 'b': 0.0149, 'c': 0.0278, 'd': 0.0425, 'e': 0.1270,
	'f': 0.0223, 'g': 0.0202, 'h': 0.0609, 'i': 0.0697, 'j': 0.0015,
	'k': 0.0077, 'l': 0.0403, 'm': 0.0241, 'n': 0.0675, 'o': 0.0751,
	'p': 0.0193, 'q': 0.0010, 'r': 0.0599, 's': 0.0633, 't': 0.0906,
	'u': 0.0276, 'v': 0.0098, 'w': 0.0236, 'x': 0.0015, 'y': 0.0197,
	'z': 0.0007,
}

func Xor(key byte, b []byte) []byte {
	result := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		result[i] = b[i] ^ key
	}
	return result
}

func RXor(key []byte, b []byte) []byte {
	diff := len(b) / len(key)
	remain := len(b) % len(key)
	newKey := make([]byte, 0, len(b))

	for i := 0; i < diff; i++ {
		newKey = append(newKey, key...)
	}
	newKey = append(newKey, key[:remain]...)

	res := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		res[i] = b[i] ^ newKey[i]
	}

	return res
}

func ChiSquaredScore(text []byte) float64 {
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

	score := 0.0
	if totalLetters == 0 {
		return math.MaxFloat64
	}

	for letter, count := range frequency {
		expected := FreqTable[letter] * float64(totalLetters)
		if expected > 0 {
			score += math.Pow(float64(count)-expected, 2) / expected
		}
	}

	return score
}

// for whoever is reading this: i hate stats
func NewChiSquared(inp []byte) (float64, float64) {
	counts := make([]int, 256)
	for _, b := range inp {
		counts[b]++
	}

	// modified frequencies
	engFreq := []float64{
		0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		0.0001, 0.0001, 0.755, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		15.843, 0.004, 0.375, 0.002, 0.008, 0.019, 0.008, 0.134,
		0.137, 0.137, 0.001, 0.001, 0.972, 0.19, 0.857, 0.017,
		0.334, 0.421, 0.246, 0.108, 0.104, 0.112, 0.103, 0.1,
		0.127, 0.237, 0.04, 0.027, 0.004, 0.003, 0.004, 0.002,
		0.0001, 0.338, 0.218, 0.326, 0.163, 0.121, 0.149, 0.133,
		0.192, 0.232, 0.107, 0.082, 0.148, 0.248, 0.134, 0.103,
		0.195, 0.012, 0.162, 0.368, 0.366, 0.077, 0.061, 0.127,
		0.009, 0.03, 0.015, 0.004, 0.0001, 0.004, 0.0001, 0.003,
		0.0001, 6.614, 1.039, 2.327, 2.934, 9.162, 1.606, 1.415,
		3.503, 5.718, 0.081, 0.461, 3.153, 1.793, 5.723, 5.565,
		1.415, 0.066, 5.036, 4.79, 6.284, 1.992, 0.759, 1.176,
		0.139, 1.162, 0.102, 0.0001, 0.002, 0.0001, 0.0001, 0.0001,
		0.06, 0.004, 0.003, 0.002, 0.001, 0.001, 0.001, 0.002,
		0.001, 0.001, 0.0001, 0.001, 0.001, 0.003, 0.0001, 0.0001,
		0.001, 0.001, 0.001, 0.031, 0.006, 0.001, 0.001, 0.001,
		0.002, 0.014, 0.001, 0.001, 0.005, 0.005, 0.001, 0.002,
		0.017, 0.007, 0.002, 0.003, 0.004, 0.002, 0.001, 0.002,
		0.002, 0.012, 0.001, 0.002, 0.001, 0.004, 0.001, 0.001,
		0.003, 0.003, 0.002, 0.005, 0.001, 0.001, 0.003, 0.001,
		0.003, 0.001, 0.002, 0.001, 0.004, 0.001, 0.002, 0.001,
		0.0001, 0.0001, 0.02, 0.047, 0.009, 0.009, 0.0001, 0.0001,
		0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.003, 0.001,
		0.004, 0.002, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.001,
		0.001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		0.005, 0.002, 0.061, 0.001, 0.0001, 0.002, 0.001, 0.001,
		0.001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
		0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001,
	}
	freqDist := make([]float64, 256)

	for i, c := range counts {
		if c > 0 {
			freqDist[i] = (float64(c) / float64(len(inp))) * 100
		}
	}

	// calculate chisqured
	score := stat.ChiSquare(freqDist, engFreq)
	df := float64(len(freqDist) - 1)

	return score, 1 - distuv.ChiSquared{K: df}.CDF(score)
}

func FindXOR(data []byte) (byte, string, float64) {
	bestScore := math.MaxFloat64
	bestChar := byte(0)
	bestResult := []byte{}
	bestProb := 0.0

	for i := 0; i <= 255; i++ {
		xored := Xor(byte(i), data)
		// cleaned := CleanText(xored)
		score, prob := NewChiSquared(xored)

		if prob >= bestProb && score > 0 {
			bestScore = score
			bestChar = byte(i)
			bestResult = xored
			bestProb = prob
		}

		/*
			score := ChiSquaredScore(cleaned)

			if score < bestScore && score > 0 {
				bestScore = score
				bestChar = byte(i)
				bestResult = cleaned
			}
		*/
	}

	/*
		fmt.Printf("Best XOR character: '%c' (0x%X)\n", bestChar, bestChar)
		fmt.Printf("XORed result: '%s'\n", string(bestResult))
		fmt.Printf("Best Score: %f\n", bestScore)
	*/

	return bestChar, string(bestResult), bestScore
}
