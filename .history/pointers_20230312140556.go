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
	i := 10
	var ptr3 *int = &i
	
	//can also be initialized without specifying data-type
	s := "hello"
	var ptr4 = &s

	//can also be intialized by short hand and not specifying var keyword
	str := "hey"
	ptr5 := &str

	fmt.Println(ptr3, ptr4, ptr5)
	fmt.Println(*ptr3, *ptr4, *ptr5) //dereferencing using *<pointer_name>

	//changing value at specific address using the pointer
	*ptr5 = "Good day"
	fmt.Println(*ptr5)

}