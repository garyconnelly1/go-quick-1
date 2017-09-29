//Author: Gary Connelly
//date 29/09/2017

package main

import (
	"fmt"
	"strings"
)

func main() {
 	var word string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &word)
	word = strings.ToLower(word)
	fmt.Println(palindromeChecker(word))
}

func palindromeChecker(s string) string {
	middle := len(s) / 2
	last := len(s) - 1
	for i := 0; i < middle; i++ {
		if s[i] != s[last-i] {
			return "NO. It's not a Palimdrome."
		}
	}
	return "YES! You've entered a Palindrome"
}