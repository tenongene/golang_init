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
	fmt.Printf("%+v \n", s)   //+v is special format specifier for structs

//by using shorthand assignment and the "new" keyword:  <variable_name> := new(<struct_name>)
	st := new(Student)        //basically a pointer of type Student
	fmt.Printf("%+v \n", st)

//by using the shorthand assignment: <variable_name> := <struct_name> { <field_name>: <value>, <field_name>: <value> }
	pupil1 := Student{
		name "Joe"
		rollNo: 15,
		score: [20]
	}

}