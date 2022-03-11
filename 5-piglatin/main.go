package main

import (
	"fmt"
	"strings"
)

const (
	pigLatinSuffix             string = "ay"
	firstLetterExceptions      string = "aeiou"
	firstLetterExceptionSuffix string = "d" + pigLatinSuffix
)

// Translate translates one or more english words into the PigLatin equlivent
func Translate(in string) string {
	var pigLatinWords []string
	englishWords := strings.Split(in, " ")

	for _, word := range englishWords {
		first := word[0:1]
		if strings.Contains(firstLetterExceptions, first) {
			pigLatinWords = append(pigLatinWords, word+firstLetterExceptionSuffix)
		} else {
			pigLatinWords = append(pigLatinWords, word[1:]+first+pigLatinSuffix)
		}
	}
	return strings.Join(pigLatinWords, " ")
}

func main() {
	fmt.Println(Translate("crnch"))
}

/*
piglatin
Instructions
Write a program that transforms a string passed as argument in its Pig Latin version.

The rules used by Pig Latin are as follows:

If a word begins with a vowel, just add "ay" to the end.
If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
If the word has no vowels, the program should print "No vowels".

If the number of arguments is different from one, the program prints nothing.

Usage
$ go run .
$ go run . pig | cat -e
igpay$
$ go run . Is | cat -e
Isay$
$ go run . crunch | cat -e
unchcray$
$ go run . crnch | cat -e
No vowels$
$ go run . something else | cat -e
$
*/
