//Author: Gary Connelly
//date 02/10/2017

package main

import (
	"fmt" 
)


func main() {
	word := "HelloWorld" // Word to reverse

	fmt.Printf("The reverse of the String '%s' is %s", word, reverseWord(word)) // Print string in reverse
}

func reverseWord(s string) string {
	var reverse string
	
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}