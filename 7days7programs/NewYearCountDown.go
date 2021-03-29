package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time  : ", currentTime)

	format := "2006-01-02 15:04:05"
    newYear,_ := time.Parse(format, "2021-12-31 12:00:00")
    fmt.Println(newYear)

    //diff := currentTime.Sub(newYear)
	diff := newYear.Sub(currentTime)
    //func Since(t Time) Duration
    //Since returns the time elapsed since t. 
    //It is shorthand for time.Now().Sub(t).
	fmt.Println("New Year 2021")
	fmt.Println("Days :",int(diff.Hours()/24))// number of days    
    fmt.Println("Hours :",diff.Hours())// number of Hours
    fmt.Println("Minutes :",diff.Minutes())// number of Minutes
    fmt.Println("Seconds :",diff.Seconds())// number of Seconds
}