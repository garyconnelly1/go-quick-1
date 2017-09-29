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
