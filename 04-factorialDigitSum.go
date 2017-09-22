//Author: Gary Connelly
//date 22/09/2017

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