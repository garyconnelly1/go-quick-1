//Author: Gary Connelly
//date 20/09/2017

//Write a program that prints the current time and date to the console.


//Adapted from https://tour.golang.org/welcome/4



// main method
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now()) //time.Now() is method for getting time.
}
