package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	first := os.Args[1]
	second := os.Args[2]

	sFirst := []rune(first)
	sSecond := []rune(second)
	result := []rune{}

	if len(args) != 3 {
		fmt.Println()
	} else {
		for f := 0; f < len(sFirst); f++ {
			for s := 0; s < len(sSecond); s++ {
				if sFirst[0] == sSecond[s] {
					sSecond = sSecond[s+1:]
					result = append(result, sFirst[0])
					sFirst = sFirst[1:]
					//fmt.Println(string(result))
				}
			}
		}
	}
	if string(result) == first {
		fmt.Println(string(result))
	}
}

/*
Write a program that takes two strings and checks whether it’s possible to
write the first string with characters from the second string,
while respecting the order in which these characters appear in the second string.
If it's possible, the program displays the string, followed by a \n, otherwise it simply displays a \n.
If the number of arguments is not 2, the program displays a \n.


faya fgvvfdxcacpolhyghbreda
faya fgvvfdxcacpolhyghbred
error rrerrrfiiljdfxjyuifrrvcoojh
"quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
*/
