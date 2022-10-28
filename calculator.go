package main

import (
	"fmt"
	//"strconv"
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

   }
}