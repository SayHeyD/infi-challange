package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func generateRandomNumber() int {
	return rand.IntN(100)
}

func checkNumber(targetNumber int, enteredNumber int) bool {
	if targetNumber > enteredNumber {
		fmt.Println("The target number is greater than the entered number")
		return false
	} else if targetNumber < enteredNumber {
		fmt.Println("The target number is lower than the entered number")
		return false
	} else {
		fmt.Println("Congratulations, you are entered the correct number!")
		return true
	}
}

func playAgain(inputReader *bufio.Reader) bool {
	fmt.Println("Do you want to play again? (Y/n)")
	for {
		userInput, _ := inputReader.ReadString('\n')
		cleanedUserInput := strings.ToLower(strings.TrimSpace(userInput))
		if cleanedUserInput == "y" || cleanedUserInput == "yes" {
			return true
		} else if cleanedUserInput == "n" || cleanedUserInput == "no" {
			return false
		} else {
			fmt.Println("Please only enter 'y', 'yes', 'n' or 'no'.")
		}
	}
}

func main() {
	playing := true

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Number guessing Game!")
	fmt.Println()
	fmt.Println("Enter a number and you will be told if the target number is greater or lower than the number you entered.")
	fmt.Println()
	fmt.Println("The Number is between 0 and 100")

	targetNumber := generateRandomNumber()

	for playing {
		fmt.Println("Guess the number: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read your input, please try again: " + err.Error())
			continue
		}
		fmt.Println(userInput)
		inputAsInt, err := strconv.Atoi(strings.TrimSpace(userInput))
		if err != nil {
			fmt.Println("Input must be a number (Integer), please try again...")
		} else {
			if checkNumber(targetNumber, inputAsInt) {
				if playAgain(reader) {
					targetNumber = generateRandomNumber()
				} else {
					playing = false
				}
			}
		}
	}

	fmt.Println("Thank you for playing!")
	os.Exit(0)
}
