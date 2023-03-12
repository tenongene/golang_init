package main 

import "fmt" 

//A pointer is a variable that holds the memory address of another variable
//Provide a way to find or change the value located at a particular memory location

//address-of-operator: preceded by "&" - gives memory address
//dereference-operator: preceded by "*" - gives the value at memory address

//declaring pointer syntax: var <pointer_name> *<data_type> (asterix not same as dereferance-operator)


func main () { 
	x := 55
	fmt.Println(&x)  //returned 0xc0000ba000
	fmt.Println(*(&x)) //dereferencing

	//declaring pointers
	var ptr1 *int
	var ptr2 *string
	fmt.Println(ptr1)
	fmt.Println(ptr2)

	//initializing pointers
	var

}