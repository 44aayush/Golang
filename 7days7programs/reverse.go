package main

import "fmt"

func reverse(str string) string {
	var rev string
	
	for i := len(str)-1; i>=0; i-- {
		rev = rev + string(str[i])
	} 
	return rev
}

func main() {
	var str string
	fmt.Println("Enter string to reverse")
	fmt.Scan(&str)

	rev := reverse(str)
	fmt.Printf("Reverse String is %s",rev)
}