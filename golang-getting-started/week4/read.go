/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter a filename: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := make([]name, 0, 3)
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		_fname := names[0]
		_lname := names[1]
		_name := name{_fname, _lname}
		s = append(s, _name)
	}

	for _, _name := range s {
		fmt.Println("fname: " + _name.fname + " lname: " + _name.lname)
	}
}
