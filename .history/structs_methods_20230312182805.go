package main 

import "fmt" 

//struct - user defined data-type, group together data elements, reference grouped values via a single variable
//struct syntax: type<struct_name>struct{ //list of fields }

//declaration
type Student struct {
	name string
	rollNo int	
	// scores []int
	// grades map[string]string
}




func main () { 

//initialization - 3 ways
//by creating a variable of data-type struct: var <variable_name><struct_name>
	var s Student
	fmt.Printf("%+v \n", s)   //+v is special format specifier for structs

//by using shorthand assignment and the "new" keyword:  <variable_name> := new(<struct_name>)
	st := new(Student)        //basically a pointer of type Student
	fmt.Printf("%+v \n", st)

//by using the shorthand assignment: <variable_name> := <struct_name> { <field_name>: <value>, <field_name>: <value> }
	pupil1 := Student{
		name: "Joe",
		rollNo: 15,
		// scores: []int[95],
		// grades: map[string]string{"Math": "A"}
	}
	fmt.Printf("%+v \n", pupil1)


	//Accessing fields of a struct using dot notation:  <variable_name

}