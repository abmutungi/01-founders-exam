package main

import "fmt"

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb - 1
		for IsPrime(n) == false {
			n--
		}
		return n
	}
	return 0
}


func main() {
	fmt.Println(FindPrevPrime(44))
	fmt.Println(FindPrevPrime(72))
}

/*
findprevprime
Instructions
Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

If there are no primes inferior to the int passed as parameter the function should return 0.

Expected function
func FindPrevPrime(nb int) int {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.FindPrevPrime(5))
	fmt.Println(piscine.FindPrevPrime(4))
}
And its output :

$ go run .
5
3
$
*/