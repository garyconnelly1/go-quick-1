//Author: Gary Connelly
//date 22/09/2017

//Write a guessing game where the user must guess a randomly generated number. After every guess the program tells the user whether their number was too high or too low. At the end, the number of tries should be printed. It counts only as one try if they input the same number multiple times consecutively.

//https://www.youtube.com/watch?v=gh1yOouqFs0

package main

//import "time"
import "fmt"
import "math/rand"

func main (){
  
  var myNum int
  myNum = rand.Intn(100) 
  //fmt.Println(myNum)
  
  count := 0
	var guess int