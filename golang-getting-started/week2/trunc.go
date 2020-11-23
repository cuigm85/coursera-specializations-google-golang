/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputedNumber string
	fmt.Printf("Enter a floating point number: ")
	_, _ = fmt.Scan(&inputedNumber)
	truncatedNumber, _ := strconv.Atoi(
		inputedNumber[0:strings.Index(inputedNumber, ".")])
	fmt.Printf("Truncated version of the floating point number is: %d\n", truncatedNumber)
}
