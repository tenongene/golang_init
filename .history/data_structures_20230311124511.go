package main
import "fmt"


func main() {

	//array syntax" " var <array name> [<size>] <data type>

	// var grades [5] int
	// var fruits [4] string

	//initialization - 3 ways
	// var grades [5] int = [3]int{10, 20, 30, 40, 50}
	// grades := [5]int{10, 20, 30, 40, 50}
	// grades := [...]int{10, 20, 30, 40, 50}

	fruits := [2]string{"apples", "bananas"}
	fmt.Println(fruits, len(fruits))
	fmt.Println(fruits[1])

	//for loop of arrays - just like js
	//for loop using "range" - range returns index and element (just like "" in js)

	for index, element := range fruits {
		fmt.Println(index,"=>", element)
	}

}