//Author: Gary Connelly
//date 22/09/2017

//Write a function that calculates the sum of the digits of the factorial of a number. n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. Then find the sum of the digits in the number 100!.

//adapted from https://gist.github.com/shockalotti/7fa310e915ee66766039

package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	
	return x * factorial(x-1)
}

func main() {
	x := int(100)
	calcFactorial := factorial(x)
	fmt.Println(calcFactorial)
}