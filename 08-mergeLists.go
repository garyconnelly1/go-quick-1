//Author: Gary Connelly
//date 29/09/2017

package main


import 
	"fmt"
	


func main () {
	 x := [] int {2,4,6}
	y := [] int {1,5,7}
  
  //displaying the lists
	fmt.Println("List 1 is: ")
	for _, v := range x {
    fmt.Println(v)
	}
  
	fmt.Println("List 2 is: ")
	for _, v := range y {
    fmt.Println(v)
	}
  
	//merge lists
	var list3 [] int = merge (x, y) //merge is the function from below
	fmt.Println("The merged list is: ")
  
  //iterate through the merged sorted list
	for _, v := range list3 {
    fmt.Println(v)
	}//end for
	fmt.Println()

}//end main


//function to merge the sorted lists
func merge (list1 [] int, list2 [] int) [] int{
  var list3 [] int = len(list1) + len(list2)
  
  //create the first half of list3 using list 1
  for i := 0; i < len(list1); i++ {
    list3[i] = list1[i]
    
  }//end for
  
  //create the second half of list3 using list 2
  for i = 0 , j := len(list1); i < len(list2); i++ , j++ { 
    list3[j] = list2[i]
    
  }//end for
  
   sort(list3)//call function that sorts the list
   return list3
  
}//end merge method

// method for sorting lists
func sort (list [] int){
   for i := 0; i < len(list) -1; i++{
     var min int = list [i]
     var minIndex int = i
     
    for j := i + 1; j < len(list); j++{
      //bringing the smallest element of the list to the front of the index
      if (list[j] < min){
        
        min = list[j]
		//change the index
        minIndex = j
        
      }//end if
      
      
    }//end inner for
    
    if (minIndex != i){
      list[minIndex] = list[i]
      list[i] = min
    }//end if
   }//end for
  
}//end sort function