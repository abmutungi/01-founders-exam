package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	count := 0

	for _, v := range args {
		_ = v
		count++
	}
	fmt.Println(count)
}
