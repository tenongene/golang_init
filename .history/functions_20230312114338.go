package main
import "fmt"


func addNumbers (a int, b int) int {
		sum := a + b
		return sum
	}

//can return multiple values 


func main() {

	//funtions syntax func <function_name>(<params>) <return_type> { <body> }
	fmt.Println(addNumbers(2,3))

}
