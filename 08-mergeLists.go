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


//function to merge the sorted lists
func merge (list1 [] int, list2 [] int) [] int{
  var list3 [] int = len(list1) + len(list2)
  
  for i := 0; i < len(list1); i++ {
    list3[i] = list1[i]
    
  }//end for
  
  for i = 0 , j := len(list1); i < len(list2); i++ , j++ { 
    list3[j] = list2[i]
    
  }//end for
  
   sort(list3)
   return list3
  
}//end merge method

// method for sorting lists
func sort (list [] int){
   for i := 0; i < len(list) -1; i++{
     var min int = list [i]
     var minIndex int = i
     
    for j := i + 1; j < len(list); j++{
      
      if (list[j] < min){
        
        min = list[j]
        minIndex = j
        
      }//end if
      
      
    }//end inner for
    
    if (minIndex != i){
      list[minIndex] = list[i]
      list[i] = min
    }//end if
   }//end for
  
}//end sort function