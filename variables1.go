package main

import "fmt"

func main() {
	//var number = 2.2
	//number  -= 10
	number := 6 //Walrus Operator
	fmt.Println(number)
	fmt.Printf("%T", number)

	abc := "Hello"
	//abc := 6
	fmt.Println("\n",abc)

	boo := true
	fmt.Println(boo)

	//Default values
	var numb uint64
	fmt.Println(numb)

	var bl bool
	fmt.Println(bl)

	def := 10 
	fmt.Printf("Hello %v %%", def%2) // %% literal % sign to print

	var out string = fmt.Sprintf("\nHello World, Number is %d", 01)
	fmt.Println(out)
}