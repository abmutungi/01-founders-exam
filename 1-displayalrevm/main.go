package main

import (
	"fmt"
	"strings"
)

func main() {
	alpha := "zyxwvutsrqponmlkjihgfedbca"

	for i, v := range alpha {
		if i%2 == 0 {
			fmt.Print(string(v))
		} else {
			fmt.Print(strings.ToUpper(string(v)))
		}
	}
	fmt.Println()
}
