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
func 




func main() {

	//funtions syntax func <function_name>(<params>) <return_type> { <body> }
	fmt.Println(addNumbers(2,3))
	fmt.Println(operations(9,6))

	sum, diff := operate(12,6)
	fmt.Println(sum, diff)

}
