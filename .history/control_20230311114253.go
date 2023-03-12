package main
import "fmt"

func main ()  {
	
	mood := "joy"
	if (mood == "happy") {
		fmt.Println(mood)
	} else {
		fmt.Println("Not equal")
	}

	//switch statements  - "fallthrough" keyword forces execution to continue (opposite of 'break' in js)
	var number int = 8
	switch number {
		case 1:
			fmt.Println("number is 1")
		case 100, 200:
			fmt.Println("number is 100 or 200")
		default:
			fmt.Println("number is neither")
	}


	//Loops
	//syntax " for initialization; condition; post"
}