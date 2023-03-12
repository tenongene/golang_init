package main

import "fmt"

// comment
func main() {

	//printing
	fmt.Println("Hello World!")

	//declaring
	name := "Bryson"
	count:= 9
	var s string = "status"
	var i int = 25
	var q bool = false
	test := true
	fmt.Println(name, count)
	fmt.Println(s, i, q)
	fmt.Println(test)

	//Formatted Printing/interpolation (format specifier)
	// %v - default, %d - intergers, %c - character, %s - plain string, %t - boolean, %f - float, %.2f - 2 decimal_float
	user := "Harry"
	role := "Engineer"
	age := 30
	fmt.Printf("%v is %d years old and is an %v", user, age, role)

	//User Input - Scanf (format specifier) - use ampersand (&) with variable
	var title string
	fmt.Print("Enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello there, ", name)

}

