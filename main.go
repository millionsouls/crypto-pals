package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"crypto-tools/encoding"
)

func main() {
	base64Cmd := &encoding.Base64Command{}
	hexCmd := &encoding.HexCommand{}

	fmt.Println("Welcome to Crypto Tools! Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) < 2 {
			fmt.Println("Invalid command. Type 'help' for usage.")
			continue
		}

		command, action := parts[0], parts[1]
		data := strings.Join(parts[2:], " ")

		switch command {
		case "base64":
			if action == "encode" {
				fmt.Println(base64Cmd.Encode(data))
			} else if action == "decode" {
				decoded, err := base64Cmd.Decode(data)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println(decoded)
				}
			}
		case "hex":
			if action == "encode" {
				fmt.Println(hexCmd.Encode(data))
			} else if action == "decode" {
				decoded, err := hexCmd.Decode(data)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println(decoded)
				}
			}
		case "help":
			fmt.Println(base64Cmd.Help())
			fmt.Println(hexCmd.Help())
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}
