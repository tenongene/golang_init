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

	//slices- extract array from reference array - mutates array - (defined as a reference to an underlying array)
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
	fmt.Println(cap(subslice))

	//declaring a slice using "make" - syntax   slice := make([]<data_type>, length, capacity)
	//usually used to create empty slice - capacity must be specified
	slice2 := make([]int,5,10)  
	fmt.Println(slice2)     //returns empty array - uninitialized      
	fmt.Println(len(slice2))        
	fmt.Println(cap(slice2))        

	//appending to slice  - syntax:   func append (<slice>[]<data_type, values) []T
	//capacity is doubled if new append exceeds capacity
	slice3 := append(subslice, 10, 20, 30)
	fmt.Println(slice3)

	//append 2 slices using a "..." operator
	slice4 := append(slice3, slice1...)
	fmt.Println(slice4)

	//copy from a slice - syntax:  func copy(dst, src []Type) - both slices must have the same data types
	dest_slice := make([]int, 3)
	num := copy(dest_slice, slice4)
	fmt.Println(dest_slice)
	fmt.Println(num)
	


	//MAPS
	// declaring syntax:   var <map_name> map[<data_type_for_key>]<data_type_for_value>  -- then initialize
	// initializing syntax: <map_name

	var codes map[string]string
}