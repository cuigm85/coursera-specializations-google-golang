/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	temp := strings.ToLower(input)
	if strings.HasPrefix(temp, "i") && strings.HasSuffix(temp, "n") && strings.Contains(temp, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
