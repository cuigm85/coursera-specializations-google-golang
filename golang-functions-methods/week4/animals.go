/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (a Bird) Eat() {
	fmt.Println("worms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	animals := make(map[string]Animal)
	for {
		fmt.Print("Enter your request (Ctrl + c to Exit)\n> ")
		scanner.Scan()
		tokens := strings.Split(scanner.Text(), " ")

		if tokens[0] == "newanimal" {
			if tokens[2] == "cow" {
				animals[tokens[1]] = new(Cow)
			} else if tokens[2] == "bird" {
				animals[tokens[1]] = new(Bird)
			} else if tokens[2] == "snake" {
				animals[tokens[1]] = new(Snake)
			} else {
				fmt.Println("Unsupported animal!")
			}
            fmt.Println("Created it!")
		} else if tokens[0] == "query" {
            if _, ok := animals[tokens[1]]; !ok {
				fmt.Println("Animal name not found!")
                continue
            }
			if tokens[2] == "eat" {
				animals[tokens[1]].Eat()
			} else if tokens[2] == "move" {
				animals[tokens[1]].Move()
			} else if tokens[2] == "speak" {
				animals[tokens[1]].Speak()
			} else {
				fmt.Println("Unsupported action!")
			}
		} else {
			fmt.Println("Unsupported request!")
			continue
		}
	}
}
