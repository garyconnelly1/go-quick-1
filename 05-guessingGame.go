//Author: Gary Connelly
//date 22/09/2017

//https://www.youtube.com/watch?v=gh1yOouqFs0

package main

import "fmt"



func main() {
	const myNumber = 5
	count := 1
	var guess int
	fmt.Println("Guess my number: ")
	fmt.Scanf("%v", &guess)
	
	while guess != myNumber {
	count++
		if guess < myNumber {
			fmt.Println("Too small, try again  ")
			fmt.Println("Guess my number: ")
			fmt.Scanf("%v", &guess)
			
		} else if guess > myNumber {
			fmt.Println("Too high, try again  ")
			fmt.Println("Guess my number: ")
			fmt.Scanf("%v", &guess)
			
		}
	}//end while
	
}