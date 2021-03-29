package main

import (
	"fmt"
)

func main() {

	var n, tmp, rem int
	var res int
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)

	tmp = n
	
	for {
		rem = tmp % 10
		res += rem*rem*rem
		tmp /= 10

		if (tmp == 0) {
			break
		}
	}

	if (res == n) {
		fmt.Printf("%d is an Armstrong number", n)
	} else {
		fmt.Printf("%d is not an Armstrong number", n)
	}
}