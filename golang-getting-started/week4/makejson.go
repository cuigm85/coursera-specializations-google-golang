/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Enter a address: ")
	scanner.Scan()
	address := scanner.Text()

	createdMap := map[string]string{"name": name, "address": address}
	json, err := json.Marshal(createdMap)
	if err == nil {
		fmt.Printf("%s\n", json)
	}

}
