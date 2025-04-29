package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type your name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("INPUT ERROR: ", err)
		return
	}
	name = strings.TrimSpace(name)

gameloop:
	for {
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

		for {

			fmt.Println("Do you want to play again?\n[Y]es\n[N]o")
			var want rune
			_, err := fmt.Scanf("%c", &want)
			if err != nil {
				fmt.Println("INPUT ERROR, try again!")
				continue
			} else {
				if want == 'n' {
					break gameloop
				} else if want == 'y' {
					break
				} else {
					continue
				}
			}

		}

	}

}
