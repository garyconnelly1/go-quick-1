//Author: Gary Connelly
//date 30/09/2017

//adapted from https://gist.github.com/abesto/3476594

package main

import (
	"fmt"
)

func main() {
  var yourNumber float64
  fmt.Println("Enter your number:")
  fmt.Scanf("%f\n", &yourNumber)
	fmt.Println(Sqrt(yourNumber))

}