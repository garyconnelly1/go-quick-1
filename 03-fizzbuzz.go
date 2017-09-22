//Author: Gary Connelly
//date 22/09/2017


package main

import (
	"fmt"
	
)
//main methof

func main() {
	//sum:= 0;
	for i := 0; i < 100; i++ {
   	// sum += i
		
	 if i % 3 == 0 {
		fmt.Println("fizz")
	 } else if i % 5 == 0 {
		fmt.Println("buzz")
	 } else if i % 15 == 0  {
		fmt.Println("fizzbuzz")
	 } else {
		fmt.Println(i)
	 }//end if else
	
  }
  
}