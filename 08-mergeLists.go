//Author: Gary Connelly
//date 29/09/2017

package main


import 
	"fmt"
	


func main () {
	 x := [] int {2,4,6}
	y := [] int {1,5,7}
  
	fmt.Println("List 1 is: ")
	for _, v := range x {
    fmt.Println(v)
	}
  
	fmt.Println("List 2 is: ")
	for _, v := range y {
    fmt.Println(v)
	}
  
	//merge lists
	var list3 [] int = merge (x, y)
	fmt.Println("The merged list is: ")
  
	for _, v := range list3 {
    fmt.Println(v)
	}//end for
	fmt.Println()

}//end main