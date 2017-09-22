//Author: Gary Connelly
//date 22/09/2017


package main

import (
	"fmt"
	
)
//main methof

func main() {
	sum:= 0;
	//basic for loop to print from 1-100
	for i := 0; i < 100; i++ {
    sum += i
	fmt.Println(sum)
  }
  
}