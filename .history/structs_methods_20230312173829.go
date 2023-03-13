package main 

import "fmt" 

//struct - user defined data-type, group together data elements, reference grouped values via a single variable
//struct syntax: type<struct_name>struct{ //list of fields }

//declaration
type Student struct {
	name string
	rollNo int	
	score []int
	grades map[string]int
}




func main () { 

//initialization - # ways
//by creating a variable of data-type struct: var <variable_name><struct_name>
	var s Student
	fmt.Printf("%")
}