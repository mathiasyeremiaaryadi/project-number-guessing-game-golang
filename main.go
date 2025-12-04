package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	MIN_NUMEBR = 1
	MAX_NUMBER = 100
)

type Difficulty struct {
	Name     string
	Attempts int
}

var scores = make(map[string][]int)

var difficulties = map[string]Difficulty{
	"1": {Name: "Easy", Attempts: 10},
	"2": {Name: "Medium", Attempts: 5},
	"3": {Name: "Hard", Attempts: 3},
}

func main() {
	for {
		ShowMenu()

		selectedDifficulty := GetDifficulty()
		answer := GenerateRandomNumber()

		fmt.Println()

		fmt.Printf("Great! You have selected the %s difficulty level.\n", selectedDifficulty.Name)
		fmt.Println("Let's start the game!")
		fmt.Println()

		isGuessed, successAttempt := RunGame(selectedDifficulty.Attempts, answer)

		if !isGuessed {
			fmt.Printf("You lost! You've used all your attempts. The correct number was %d.\n", answer)
		} else {
			scores[selectedDifficulty.Name] = append(scores[selectedDifficulty.Name], successAttempt)
		}

		fmt.Println()

		if !ConfirmPlayAgain() {
			fmt.Println(GetTopScore(scores))
			break
		}

		fmt.Println()
	}
}

func ShowMenu() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println()
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println()
}

func GetDifficulty() Difficulty {
	fmt.Print("Enter your choice: ")
	var choice string
	fmt.Scan(&choice)

	if difficulty, exist := difficulties[choice]; exist {
		return difficulty
	}

	fmt.Println("Invalid choice! Defaulting to Medium difficulty.")
	return difficulties["2"]
}

func GenerateRandomNumber() int {
	return rand.Intn(MAX_NUMBER-MIN_NUMEBR+1) + MIN_NUMEBR
}

func RunGame(attempts int, answer int) (bool, int) {
	start := time.Now()
	for attempt := 1; attempt <= attempts; attempt++ {
		if attempts-3 == attempt {
			fmt.Println(GetHint(answer))
		}

		var guess int
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess == answer {
			elapsed := time.Since(start).Seconds()
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts and it took %.2f seconds.\n", attempt, elapsed)
			return true, attempt
		}

		if guess < answer {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
			fmt.Println()
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
			fmt.Println()
		}
	}

	return false, 0
}

func GetHint(answer int) string {
	fmt.Println("You have 3 attempts left! Here is the hint:")

	numberType := "odd"
	if answer%2 == 0 {
		numberType = "even"
	}

	answerConverted := strconv.Itoa(answer)
	return fmt.Sprintf("The number starts with %c and it's a %s number.\n", answerConverted[0], numberType)
}

func ConfirmPlayAgain() bool {
	fmt.Print("Do you want to play again? (yes/no): ")

	var playAgain string
	fmt.Scan(&playAgain)

	fmt.Println()

	return playAgain == "yes"
}

func GetTopScore(scoreList map[string][]int) string {
	var topDifficulty string
	var topAttempt int = 0

	for difficulty, attempts := range scoreList {
		for _, attempt := range attempts {
			if topAttempt == 0 || attempt < topAttempt {
				topAttempt = attempt
				topDifficulty = difficulty
			}
		}
	}

	if topAttempt == 0 {
		return "No scores available."
	}

	return fmt.Sprintf("Your high score is %d attempts on %s difficulty.", topAttempt, topDifficulty)
}
