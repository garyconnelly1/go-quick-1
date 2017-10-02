//Author: Gary Connelly
//date 29/09/2017


//Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.

package main

import (
	"fmt"
	"strings"
)

//main method
func main() {
 	var word string
	fmt.Println("Enter string:")
	//scanf to take in user input
	fmt.Scanf("%s\n", &word)
	//convert all letters to lower case
	word = strings.ToLower(word)
	//call the palindromeChecker method
	fmt.Println(palindromeChecker(word))
}

//method to check if the string is a palindrome
func palindromeChecker(s string) string {
	middle := len(s) / 2
	last := len(s) - 1
	//for each letter up to the middle of the word, check if its the same as the corresponding letter from the end of the string
	for i := 0; i < middle; i++ {
		if s[i] != s[last-i] {
			return "NO. It's not a Palimdrome."
		}
	}
	return "YES! You've entered a Palindrome"
}