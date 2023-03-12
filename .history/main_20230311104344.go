package main

import (
	"fmt"
	"reflect"
)

// comment
func main() {

	//printing
	fmt.Println("Hello World!")

	//declaring
	name := "Bryson"
	// count:= 9
	var s string = "status"
	var i int = 25
	var q bool = false
	test := true
	fmt.Println(name)
	fmt.Println(s, i, q)
	fmt.Println(test)

	//Formatted Printing/interpolation (format specifier)
	// %v - default, %d - intergers, %c - character, %s - plain string, %t - boolean, %f - float, %.2f - 2 decimal_float
	user := "Harry"
	role := "Engineer"
	age := 30
	fmt.Printf("%v is %d years old and is an %v \n", user, age, role)

	//User Input - Scanf (format specifier) - use ampersand (&) with variable
	// var title string
	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &title)
	// fmt.Println("Hello there,", title)

	//fmt.Scanf returns argument count and err when function is run
	var a string
	var b int
	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count: ", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	
	//getting type of variable - use "reflect.TypeOf" method

	fmt.PrintLn("")

}

