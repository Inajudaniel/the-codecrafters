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
	fmt.Println("✦ CLI Calculator ✦ (type 'help')")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		parts := strings.Fields(strings.TrimSpace(input))
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		if cmd == "quit" {
			fmt.Println("Goodbye! ✦")
			return
		}
		if cmd == "help" {
			fmt.Println("Commands: add a b | sub a b | mul a b | div a b | quit | help")
			continue
		}
		if len(parts) != 3 {
			fmt.Printf("Error: '%s' needs 2 numbers.\n", cmd)
			continue
		}

		a, err1 := strconv.Atoi(parts[1])
		b, err2 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Error: arguments must be integers.")
			continue
		}

		switch cmd {
		case "add":
			fmt.Println("✦ Result:", a+b)
		case "sub":
			fmt.Println("✦ Result:", a-b)
		case "mul":
			fmt.Println("✦ Result:", a*b)
		case "div":
			if b == 0 {
				fmt.Println("Error: division by zero.")
			} else {
				fmt.Println("✦ Result:", a/b)
			}
		default:
			fmt.Printf("Error: unknown command '%s'. Type 'help'.\n", cmd)
		}
	}
}
