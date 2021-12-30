package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Printf("Welcome to bacon's random number generator!\n")
	var inputFirst string
	fmt.Scanln(&inputFirst)
	if inputFirst == "Number" {
		numberGenerator()
	} else if inputFirst == "Kill" {
		killScript()
	}
}

func numberGenerator() {
	fmt.Println("Please enter a number and the generator will generate a random from 0 to your number:")
	var random int
	fmt.Scanln(&random)
	fmt.Printf("Here is your random number:\n")
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(random))
	time.Sleep(1 * time.Second)
	fmt.Println("Would you like to play again?\n" + "If so please type Y or N to continue.")
	var playAgain string
	fmt.Scanln(&playAgain)
	if playAgain == "Y" {
		numberGenerator()
	} else if playAgain == "N" {
		main()
	}
}

func killScript() {
	var input string
	fmt.Println("Would you like to exit the script?")
	fmt.Scanln(&input)
	if input == "Y" {
		killGame()
	} else {
		main()
	}
}

func killGame() {
	fmt.Println("Script exited successfully!")
	os.Exit(1)
}
