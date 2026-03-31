package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Go Number Base Converter ---")
	fmt.Println("Usage: <number> <base> (e.g., '10 dec', 'FF hex', '1010 bin')")
	fmt.Println("Type 'quit' to exit.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "quit" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid format. Usage: <number> <base>")
			continue
		}

		numberStr := parts[0]
		baseStr := strings.ToLower(parts[1])

		switch baseStr {
		case "dec":
			val, err := strconv.ParseInt(numberStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid decimal:", err)
				continue
			}
			fmt.Printf("✦ Binary:  %b\n", val)
			fmt.Printf("✦ Hex:     %X\n", val)

		case "hex":
			val, err := strconv.ParseInt(numberStr, 16, 64)
			if err != nil {
				fmt.Println("Invalid hex:", err)
				continue
			}
			fmt.Printf("✦ Decimal: %d\n", val)

		case "bin":
			if !isValidBinary(numberStr) {
				fmt.Println("Invalid binary: contains non-binary digits")
				continue
			}
			val, err := strconv.ParseInt(numberStr, 2, 64)
			if err != nil {
				fmt.Println("Invalid binary:", err)
				continue
			}
			fmt.Printf("✦ Decimal: %d\n", val)

		default:
			fmt.Println("Unknown base. Use: dec, hex, or bin")
		}
	}
}

func isValidBinary(s string) bool {
	for _, r := range s {
		if r != '0' && r != '1' {
			return false
		}
	}
	return true
}
