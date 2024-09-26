package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

/*
determine keysize by finding top 3 lowest hamming distances
breaking ciphertext into keysize length blocks
transpose blocks
*/

func findKeySize(data []byte) int {
	bestKeySize := 0
	bestDistance := math.MaxFloat64

	for keySize := 2; keySize < 41; keySize++ {
		if len(data) < 2*keySize {
			break
		}

		firstBlock := data[:keySize]
		secondBlock := data[keySize : 2*keySize]

		distance := hamDis(firstBlock, secondBlock)
		normalizedDistance := float64(distance) / float64(keySize)

		if normalizedDistance < bestDistance {
			bestDistance = normalizedDistance
			bestKeySize = keySize
		}
	}

	return bestKeySize
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

	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	keysize := findKeySize(data)
	fmt.Println(keysize)

}
