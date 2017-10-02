//Author: Gary Connelly
//date 22/09/2017

//Write a function that returns the largest and smallest elements in a list.

// inspired by https://play.golang.org/p/vhHmjhOMEo

package main

import "fmt"

//main method
func main() {

//create list of random numbers
	x := []int{
	 26, 42, 17, 10, 4, 67, 98, 31, 82, 56,
	}
	
	//initialise the smallest and biggest to the first element in the list
	smallest, biggest := x[0], x[0]
	//for loop to iterate through list x
	for _, v := range x { // * Range - method in go for iterating through the contents of an array
		//if statements to compare the size of the element in the list
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	
	//outputs the smallest and biggest numbers
	fmt.Println("The biggest number is ", biggest)
	fmt.Println("The smallest number is ", smallest)
}