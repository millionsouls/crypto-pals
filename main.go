package main

import (
	"bufio"
	"crypto-tools/set1"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("set1/data/4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, dec := set1.AttemptSingleByteXor([]byte(scanner.Text()))
		fmt.Println(key)
		fmt.Println(dec)
	}
}
