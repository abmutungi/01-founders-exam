package main

import (
	"fmt"
	"os"
	"strconv"
)

func validOp(sum string) bool {
	op := []string{"+", "-", "*", "/", "%"}
	for _, val := range op {
		if val == sum {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	first, _ := strconv.Atoi(args[0])
	second, _ := strconv.Atoi(args[2])

	if len(args) != 3 {
		fmt.Print()
	} else {
		if validOp(args[1]) == false {
			fmt.Println(0)
		} else {
			if args[1] == "%" && second == 0 {
				fmt.Print("No Modulo by 0\n")
			} else if args[1] == "/" && second == 0 {
				fmt.Print("No division by 0\n")
			} else if args[1] == "+" {
				fmt.Println(first + second)
			} else if args[1] == "-" {
				fmt.Println(first - second)
			} else if args[1] == "*" {
				fmt.Println(first * second)
			} else if args[1] == "/" {
				fmt.Println(first / second)
			} else {
				fmt.Println(first % second)
			}
		}
	}
}

/*
doop
Instructions
Write a program that is called doop.

The program has to be used with three arguments:

A value
An operator, one of : +, -, /, *, %
Another value
In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

Usage
$ go run .
$ go run . 1 + 1 | cat -e
2$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0$
$ go run . 1 % 0 | cat -e
No modulo by 0$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
$
*/
