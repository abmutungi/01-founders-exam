package main

import "fmt"


func main() {
	for i := 9; i > 1; i-- {
		fmt.Print(rune(i))
	}
	fmt.Println('\n')
}