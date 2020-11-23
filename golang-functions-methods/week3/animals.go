/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.
*/

package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

type Animal struct {
    eat string
    move string
    speak string
}

func (a *Animal) Init(eat, move, speak string) {
    a.eat = eat
    a.move = move
    a.speak = speak
}

func (a *Animal) Eat() {
    fmt.Println(a.eat)
}
func (a *Animal) Move() {
    fmt.Println(a.move)
}
func (a *Animal) Speak() {
    fmt.Println(a.speak)
}

func main() {
    cow := new(Animal)
    cow.Init("grass", "walk", "moo")
    bird := new(Animal)
    bird.Init("worms", "fly", "peep")
    snake := new(Animal)
    snake.Init("mice", "slither", "hsss")

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Enter your request (Ctrl + c to Exit)\n> ")
        scanner.Scan()
        request := strings.Split(scanner.Text(), " ")
        animal := request[0]
        command := request[1]

        var this *Animal
        if animal == "cow" {
            this = cow
        } else if animal == "bird" {
            this = bird
        } else if animal == "snake" {
            this = snake
        } else {
            fmt.Println("The first string is the name of an animal, either “cow”, “bird”, or “snake”")
            continue
        }

        if command == "eat" {
            this.Eat()
        } else if command == "move" {
            this.Move()
        } else if command == "speak" {
            this.Speak()
        } else {
            fmt.Println("The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”")
            continue
        }
    }
}
