//Author: Gary Connelly
//date 02/10/2017

// Write a function to reverse a string in Go.

package main

//import IO functionality
import (
	"fmt" 
)


func main() {
//word could be any word at all. 
	word := "HelloWorld" // Word to reverse

	fmt.Printf("The word '%s' is %s" spelt backwards!, word, reverseWord(word)) // Print string in reverse
}

//reverse function
func reverseWord(s string) string {
	var reverse string
	
	//for i starting at the end of the word and parsing each letter into the variable reverse
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}