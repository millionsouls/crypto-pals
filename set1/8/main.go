package main

import (
	"bufio"
	"crypto-pals/util"
	"fmt"
	"os"
)

func detectECB(data []byte, size int) (map[string]int, float64) {
	chunks := util.Chunkify(data, size)
	chunkFreq := make(map[string]int)
	repeats := 0.0

	for _, chunk := range chunks {
		if _, ok := chunkFreq[string(chunk)]; ok {
			chunkFreq[string(chunk)]++
			repeats++
		} else {
			chunkFreq[string(chunk)] = 1
		}
	}

	return chunkFreq, repeats
}

func main() {
	data, err := os.Open("data.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	for line := 0; scanner.Scan(); line++ {
		text := scanner.Text()
		_, score := detectECB(util.DecodeHex(text), 16)

		fmt.Println(score)
	}

	defer data.Close()
}
