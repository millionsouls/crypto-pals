package main

// Detect single XOR
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"crypto-pals/util"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var bestResult string
	var bestKey byte
	bestScore := math.MaxFloat64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		key, result, score := util.FindXOR(util.DecodeHex(line))
		if score < bestScore && score != 0 {
			bestScore = score
			bestResult = result
			bestKey = key
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Best Result: %s\n", bestResult)
	fmt.Printf("Best Score: %f\n", bestScore)
	fmt.Printf("Best Key: %s\n", string(bestKey))
}
