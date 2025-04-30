package game

import (
	"math/rand"
	"time"
)

const RandNUm = 2

func MakeRandomNumber() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(RandNUm) + 1
}

func CheckSide(guess int, secretNumber int) rune {
	if guess < secretNumber {
		return '<'
	} else {
		return '>'
	}
}

func CheckGuess(guess int, secretNumber int) bool {
	if guess == secretNumber {
		return true
	} else if guess < secretNumber {
		return false
	} else {
		return false
	}
}
