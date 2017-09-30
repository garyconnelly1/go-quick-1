//Author: Gary Connelly
//date 30/09/2017

//adapted from https://gist.github.com/abesto/3476594

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
	return z - (z*z-x)/(2*z)
}

func Sqrt(x float64) float64 {
	z := zNext(2, x)
	for zn, delta := z, z; delta > 0.00001; z = zn {
		zn = zNext(z, x)
		delta = z - zn
	}
	return z
}