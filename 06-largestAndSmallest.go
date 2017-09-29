//Author: Gary Connelly
//date 22/09/2017

// inspired by https://play.golang.org/p/vhHmjhOMEo

package main

import "fmt"

func main() {

//create list of random numbers
	x := []int{
	 26, 42, 17, 10, 4, 67, 98, 31, 82, 56,
	}
	
	smallest, biggest := x[0], x[0]
	for _, v := range x {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	
	fmt.Println("The biggest number is ", biggest)
	fmt.Println("The smallest number is ", smallest)
}