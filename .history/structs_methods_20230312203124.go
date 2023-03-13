package main 

import "fmt" 

//struct - user defined data-type, group together data elements, reference grouped values via a single variable
//struct syntax: type<struct_name>struct{ //list of fields }

//declaring a struct
type Student struct {
	name string
	rollNo int	
	scores []int
	// grades map[string]string
}

type Circle struct {
	radius float64
	area float64
}
//creating a method for Circle struct
func (c *Circle) calcArea() {
	c.area = 3.14 * c.radius * c.radius
}

//declaring an interface
type shape interface {
	area() float64
	perim
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
		scores: []int{95},
		// grades: map[string]string{"Math": "A"}
	}
	fmt.Printf("%+v \n", pupil1)


	//Accessing fields of a struct using dot notation:  <variable_name>.<field_name>
	fmt.Println(pupil1.name)  //Joe


	//METHODS
	//A Method adds an extra parameters to a function after `func`` keyword - accepts single argument ("receiver")
	//Syntax:  func(<receiver>) <method_name> (<parameters>)<return_params> { //code }

	//creating a method for Circle struct
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v", c)
	

	//METHOD SETS
	//Set of methods available to a data type - encapsulate functionality


	//INTERFACES
	//Specifies a method set - a way to introduce modularity
	//Blueprint for method set.
	//Describe all methods of a method set by providing the function signature for each method
	//Specify method sets but do not implement them
	//Syntax: type <interface_name> interface { //method signatures }

	//Implementing an interface: - a type implements an interface by implementing its methods
	//interfaces are implemented implicitly - no explicit keyword



}