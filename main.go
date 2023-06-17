package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Welcome to Bacon's Random Number Generator!")
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Generate a random number")
		fmt.Println("2. Exit the script")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			numberGenerator()
		case 2:
			fmt.Println("Script exited successfully!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func numberGenerator() {
	fmt.Println("Please enter a number. The generator will generate a random number from 0 to your number:")

	var random int
	_, err := fmt.Scanln(&random)
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return
	}

	fmt.Println("Here is your random number:")
	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(random))
	time.Sleep(1 * time.Second)

	for {
		fmt.Println("Would you like to play again? (Y/N)")

		var playAgain string
		_, err := fmt.Scanln(&playAgain)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if playAgain == "Y" {
			break
		} else if playAgain == "N" {
			return
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Script interrupted!")
		os.Exit(0)
	}()
}

func init() {
	setupSignalHandler()
}
