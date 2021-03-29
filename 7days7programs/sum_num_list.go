//Check if number present inside the list then add the list else subtract
package main

import (
	"fmt"
)


func check_num(list []int, n int ) bool {
	for i := range list {
		if i == n {
			return true
		}
	}
	return false
}

func sum_list(num_list []int) int {
	s := 0
	for _,i := range num_list {
		s += i
	}
	return s
}

func sub_list(num_list []int) int {
	s := 0
	for _,i := range num_list {
		s -= i
	}
	return s
}

func main() {
	var n int
	var checkNum bool
	list := []int{1,2,3,4,5} 
	list1 := list
	fmt.Println("Enter a number to search")
	fmt.Scan(&n)

	checkNum = check_num(list,n)
	if checkNum == true {
		fmt.Println("Number present inside the list")
		sum := sum_list(list1)
		fmt.Printf("Sum of list is %d", sum)
	} else {
		fmt.Println("Number not present inside the list")
		sub := sub_list(list1)
		fmt.Printf("Subtarction of list is %d", sub)
	}
}
