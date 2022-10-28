package main

import (
	"fmt"
)

 func main(){
	var number1 float32
	var number2 float32
	var operator string
	
   fmt.Scanln(&number1, &operator, &number2)

   switch operator{
	   case "+":
		   fmt.Printf("%f %s %f = %f ", number1, operator, number2, number1 + number2)
	   case "-":
		   fmt.Printf("%f %s %f = %f ", number1, operator, number2, number1 - number2)
	   case "*":
		   fmt.Printf("%f %s %f = %f ", number1, operator, number2, number1 * number2)
	   case "/":
		   if (number2 == 0.0) {
		   fmt.Println("Divide by Zero situation")	
		   }else{
			   fmt.Printf("%f %s %f = %f ", number1, operator, number2, number1 / number2)
		   }
	   default:
		   fmt.Println("Error this is not number")
   }
}