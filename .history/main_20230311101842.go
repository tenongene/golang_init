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
	// %v - default, %d - intergers, %c - character, %s - pl
	user := "Harry"
	role := "Engineer"
	age := 30
	fmt.Printf("%v is %d years old and is an %v", user, age, role)


}

