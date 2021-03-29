package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
    lowerCharSet   = "abcdedfghijklmnopqrst"
    upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet = "!@#$%&*"
    numberSet      = "0123456789"
    allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
    rand.Seed(time.Now().Unix())
    var minSpecialChar, minNum, minUpperCase, passwordLength int
    fmt.Println("Enter minimum number of Special Character: ")
    fmt.Scan(&minSpecialChar)
    fmt.Println("Enter minimum number of Digits: ")
    fmt.Scan(&minNum)
    fmt.Println("Enter minimum number of UpperCase Letter: ")
    fmt.Scan(&minUpperCase)
    fmt.Println("Enter the required Password length: ")
    fmt.Scan(&passwordLength)
    // minSpecialChar := 1
    // minNum := 1
    // minUpperCase := 1
    // passwordLength := 8
    password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
    fmt.Println(password)

}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
    var password strings.Builder

    //Set special character
    for i := 0; i < minSpecialChar; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < minNum; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < minUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}