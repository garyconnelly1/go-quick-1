//Author: Gary Connelly
//date 30/09/2017

//adapted from https://gist.github.com/abesto/3476594

/*
Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.

z_next = z - ((z*z - x) / (2 * z))
This is repeated while the values of z_next and z are different. After each iteration the value of z is replaced with that of z_next.

Run a few tests to determine how close you are to the math.Sqrt value?

Hint: to declare and initialize a floating point value, give it floating-point syntax or use a conversion:

z := float64(1)
z := 1.0

*/

package main

import (
	"fmt"
)


//main method
func main() {
  var yourNumber float64
  //take in user input
  fmt.Println("Enter your number:")
  fmt.Scanf("%f\n", &yourNumber)
  //call square root method for the input number
	fmt.Println(Sqrt(yourNumber))

}

// Newton's method
func zNext(z, x float64) float64 {
	//newtons formula given in question
	return z - (z*z-x)/(2*z)
}

//method for calculate square root of any number
func Sqrt(x float64) float64 {
	//1 being the starting point for the formula to run 
	z := zNext(1, x)
	//delta being current number for z minus the next value for z
	for zn, delta := z, z; delta > 0.00001; z = zn {
		zn = zNext(z, x)
		delta = z - zn
	}
	return z
}