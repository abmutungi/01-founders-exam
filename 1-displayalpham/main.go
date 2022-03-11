package main

import (
	"fmt"
	"strings"
)

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	for i, v := range alpha {
		if i%2 == 0 {
			fmt.Print(string(v))
		} else {
			fmt.Print(strings.ToUpper(string(v)))
		}
	}
	fmt.Println()
}

/*func ToUpper(s string) string {
	a := []rune(s)

	for i := 0; i < len(s); i++ {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] = a[i] - 32
		}
	}
	return string(a)
}
*/
