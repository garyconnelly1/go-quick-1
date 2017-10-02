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
	
		//more variables
	var previousGuess int
	numFound := false
	
	//for loop for number Guessing
	for numFound != true {
	  
	  
	  fmt.Println("Guess my number:(guess must be between 1 and 100) ")
	  fmt.Scanf("%v", &guess)
	  
	  
	  if guess == previousGuess {
	    fmt.Println("That was the same as your Last guess! Try again")
	    count --
	  }//end first if
	  
	  if guess > myNum {
	    fmt.Println("To big! Please try again...")
	    count ++
	  }//end if
	  
	  if guess < myNum {
	     fmt.Println("To small! Please try again...")
	    count ++
	    
	  }//end if
	  
	  if guess == myNum {
	    numFound = true
	  }//end if
	  
	  previousGuess = guess
	}//end for loop 
	
 fmt.Println("Correct! My number was : %v ", myNum)
 fmt.Println("And it only took you %v tries!", count)
 
 
}//end main
