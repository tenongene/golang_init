package main
import "fmt"


func addNumbers (a int, b int) int {
		sum := a + b
		return sum
	}

//can return multiple values in a function (specify return type for each value)
func operations (a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

//can specify return parameter and type in the function definition, with just specifying 'return' in the body
func operate(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

//variadic function (variable number of args) - final parameter_type preceded by ellipsis "..." - must be last
//func sumNumbers(numbers ...int) int {}  // accepts any number of arguments of type int (stored in a slice called 'numbers)
//func sumNumbers(str string, numbers ...int) // variadic parameters must be last argument
func printSubjects(student string, subjects ...string) {
	fmt.Println ("Hello ", student, ", here are your subjects: ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

//blank identifiers are use to ignore return values from functions
func f() (int, int) {
	return 20, 35
}



func main() {

	//funtions syntax func <function_name>(<params>) <return_type> { <body> }
	fmt.Println(addNumbers(2,3))
	fmt.Println(operations(9,6))

	sum, diff := operate(12,6)
	fmt.Println(sum, diff)

	printSubjects("Joe", "Math", "Physics", "Civics")

	//blank identifiers are use to ignore return values from functions
	// a, b := f()
	// fmt.Println(a,b)
	_, b := f()
	fmt.Println(b)


//Anonymous functions - without any identifiersto refer to it - short term use- un-named functionality
	x := func (c int, d int) int {
		return c*d
	}
	fmt.Println(x(20,30))
	fmt.Printf("%T \n", x)
	//direct execution with no need to save to variable
	y := func (e int, f int) int {
		return e*f
	}(20,30)
	fmt.Println(y)

	fmt.Scanf("%d", "&value")


}

