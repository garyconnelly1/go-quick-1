//Author: Gary Connelly
//date 22/09/2017

package main

import "fmt"

const myNumber = 5

func main() {
	var guess int
	fmt.Println("Guess my number: ")
	fmt.Scanf("%v", &guess)
	
	while guess != myNumber {
		if guess < myNumber {
			fmt.Println("Too small, try again  ")
			fmt.Println("Guess my number: ")
			fmt.Scanf("%v", &guess)
			count++
		} else {
			fmt.Println("Too high, try again  ")
			fmt.Println("Guess my number: ")
			fmt.Scanf("%v", &guess)
			count++
		}
	}//end while
}