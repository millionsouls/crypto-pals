package main

import (
	"fmt"
	"os"
	"sort"

	"crypto-pals/util"
)

/*
determine keysize by finding top 3 lowest hamming distances
breaking ciphertext into keysize length blocks
transpose blocks
*/

type KeySize struct {
	size int
	hd   float64
}

func findKeySize(data []byte) int {
	keySizes := make([]KeySize, 39)

	for keySize := 2; keySize < 41; keySize++ {
		firstBlock := data[:keySize]
		secondBlock := data[keySize : 2*keySize]

		dis := hamDis(firstBlock, secondBlock)
		hd := float64(dis) / float64(keySize)

		keySizes[keySize-2] = KeySize{keySize, hd}

		/*
			if normalizedDistance < bestDistance {
				bestDistance = normalizedDistance
				bestKeySize = keySize
			}
		*/
	}

	sort.Slice(keySizes, func(i, j int) bool { return keySizes[i].hd < keySizes[j].hd })

	return keySizes[0].size
}

// what the hell
func chunkify(data []byte, size int) [][]byte {
	if size < 1 {
		panic("Invalid chunksize")
	}

	var chunk [][]byte

	for i := 0; i < len(data); i += size {

		end := i + size
		if end > len(data) {
			end = len(data)
		}
		nChunk := make([]byte, len(data[i:end]))

		copy(nChunk, data[i:end])

		chunk = append(chunk, nChunk)
	}

	return chunk
}

// ...
func computeKey(data []byte, key int) ([]byte, []byte) {
	chunks := chunkify(data, key)
	transpose := make([][]byte, key)

	for i := 0; i < len(chunks); i++ {
		for j := 0; j < len(chunks[i]); j++ {
			transpose[j] = append(transpose[j], chunks[i][j])
		}
	}

	keys := make([]byte, key)
	for i, tbl := range transpose {
		k, _, _ := util.FindXOR(string(tbl))
		keys[i] = k
	}

	dec := util.RXor(keys, data) //
	return dec, keys
}

func hamDis(one []byte, two []byte) int { // to util
	oneLen := len(one)
	twoLen := len(two)
	diff := 0

	if oneLen != twoLen {
		panic("Inputs of different lengths")
	}

	for i, j := range one {
		nb := two[i]
		for k := 1; k < 129; k = 2 * k {
			if (j & byte(k)) != (nb & byte(k)) {
				diff++
			}
		}
	}

	return diff
}

func main() {
	one := "this is a test"
	two := "wokka wokka!!!"

	hamming := hamDis([]byte(one), []byte(two))
	fmt.Println(hamming)

	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	keysize := findKeySize(data)
	fmt.Println(keysize)

	// keySize := findKeySize(data)
	dec, _ := computeKey(data, 29)

	fmt.Println(string(dec))
}
