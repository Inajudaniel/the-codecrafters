package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToDecimal(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {

		switch words[i] {
		case "(hex)":

			value, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")

}

func main() {
	fmt.Println(HexToDecimal("we have 1E (hex) files added"))
}
