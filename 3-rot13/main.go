package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]

	for i := 0; i < 1; i++ {
		fmt.Println(Rot13(args))
	}
}

func Rot13(str string) string {
	arg := []rune(str)
	for i := 0; i < len(arg); i++ {
		if arg[i] >= 'A' && arg[i] <= 'N' {
			arg[i] = arg[i] + 13
		} else if arg[i] >= 'N' && arg[i] <= 'Z' {
			arg[i] = arg[i] - 13
		} else if arg[i] >= 'a' && arg[i] <= 'n' {
			arg[i] = arg[i] + 13
		} else if arg[i] >= 'n' && arg[i] <= 'z' {
			arg[i] = arg[i] - 13
		}
	}
	return string(arg)
}

/*
rot13
Instructions
Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

The output will be followed by a newline ('\n').

If the number of arguments is different from 1, the program displays nothing.

Usage
$ go run . "abc"
nop
$ go run . "hello there"
uryyb gurer
$ go run . "HELLO, HELP"
URYYB, URYC
$ go run .
$
*/
