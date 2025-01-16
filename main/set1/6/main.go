package main

// Breaking repeating-key XOR
import (
	"fmt"
	"os"
	"sort"

	"crypto-pals/lib/util"
)

/*
determine keysize by finding top 3 lowest hamming distances
breaking ciphertext into keysize length blocks
transpose blocks
*/

type Keysize struct {
	size int
	hd   float64
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

func findKeySize(data []byte) int {
	keySizes := make([]Keysize, 39)

	for ks := 2; ks < 41; ks++ {
		avgDist := 0.0
		iter := 0.0

		for i := 0; i+2*ks < len(data); i += ks {
			lb, ub := data[i:i+ks], data[i+ks:i+2*ks]
			avgDist += float64(hamDis(lb, ub)) / float64(ks)
			iter++
		}

		nhd := avgDist / iter
		keySizes[ks-2] = Keysize{ks, nhd}
	}

	sort.Slice(keySizes, func(i, j int) bool { return keySizes[i].hd < keySizes[j].hd })

	return keySizes[0].size
}

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

		cchunk := make([]byte, len(data[i:end]))
		copy(cchunk, data[i:end])
		chunk = append(chunk, cchunk)
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
		k, _, _ := util.FindXOR(tbl)
		keys[i] = k
	}

	dec, _ := util.Xor(keys, data) //
	return dec, keys
}

func main() {
	/*
		one := "this is a test"
		two := "wokka wokka!!!"

		hamming := hamDis([]byte(one), []byte(two))
		fmt.Println(hamming)
	*/

	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	ddata := util.DecodeB64(string(data))
	keysize := findKeySize(ddata)

	// fmt.Println(keysize)

	text, key := computeKey(ddata, keysize)

	fmt.Println(string(text))
	fmt.Println("Key: ", string(key))
}
