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

	//multidimensional arrays/ array of arrays
	numbers := [3][2]int{{2,4}, {3,5}, {4,9}} 
	fmt.Println(numbers[0][1])

	//for loop of arrays - just like js
	//for loop using "range" - range returns index and element (just like "enumerate" in python)

	for index, element := range fruits {
		fmt.Println(index, "=>", element)
	}

	//slices- extract array from reference array - mutates array
	//new slice is declared with a slice name variable
	//capacity "cap()" - total number array can hold

	//declaring a slice imperatively
	slice := []int{10,20,30}  //note here what makes this a slice is length not defined
	fmt.Println(slice)

	//delare from another array  - syntax: array[start_index:end_index]
	arr := [10]int{12,15,17,25,29,32,37,42,48,50}
	slice1 := arr[1:8]
	fmt.Println(slice1)
	//can create slice of a slice
	subslice := slice1[0:3]
	fmt.Println(subslice)

	//declaring a slice using "make" - syntax   slice := make([]<data_type>, length, capacity)
	//usually used to create empty slice - capacity must be specified
	slice2 := make([]int,5,10)  
	fmt.P        

	
}