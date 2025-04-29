package main

import (
	"fmt"
	"guessing-game/game"
)

func main() {

	for {
		game.PlayGame(game.DecideName())
		if !game.PlayAgain() {
			fmt.Println("Thanks for playing!")
			break
		}

	}
}
