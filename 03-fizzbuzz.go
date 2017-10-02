//Author: Gary Connelly
//date 22/09/2017

//Write a program that prints the numbers from 1 to 100, each on a new line, to the console, except for the following conditions. For multiples of three print Fizz instead of the number, and for multiples of five print Buzz. For numbers that are multiples of both three and five print FizzBuzz.

//adapted from https://data-representation.github.io/notes/go.html

package main

import (
	"fmt"
	
)
//main methof

func main() {
	
	for i := 0; i < 100; i++ {
	
	// if else conditions for fizz and buzz 
	 if i % 3 == 0 {
		fmt.Println("fizz")
	 } else if i % 5 == 0 {
		fmt.Println("buzz")
	 } else if i % 15 == 0  {
		fmt.Println("fizzbuzz")
	 } else {
		fmt.Println(i)
	 }//end if else
	
  }//end for
  
}//end main method