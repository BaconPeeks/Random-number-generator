package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Welcome to bacon's random number generator!\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Please enter a number and the generator will generate a random from 0 to your number:")
	var random int
	fmt.Scanln(&random)
	fmt.Printf("Here is your random number:\n")
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(random))
	time.Sleep(1 * time.Second)
}
