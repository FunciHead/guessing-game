package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PlayGame(name string) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	fmt.Println("Hello:", name, " you need to guess a number between 1 and 1000 you have ten attempts ")
	secretNumber := r.Intn(1000) + 1
	var winOrLose bool = false
	var attempts int = 10

	for attempts > 0 {
		fmt.Println("Attempts left: ", attempts)
		fmt.Print("Make your guess: ")

		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("INPUT ERROR!\nTry Again")
			continue
		} else {
			if guess == secretNumber {
				attempts = 0
				fmt.Println("YES!! YOU GUESSED RIGHT THE SECRETE NUMBER REALLY IS: ", guess)
			} else if guess < secretNumber {
				fmt.Println("HINT: The secret number is bigger than :", guess)
				attempts--
			} else {
				fmt.Println("HINT: The secret number is smaller than :", guess)
				attempts--
			}
		}

	}

	if winOrLose {
		fmt.Println("Congratulations! You won want to play again?")
	} else {
		fmt.Println("You lost! The number is: ", secretNumber)
	}

}
