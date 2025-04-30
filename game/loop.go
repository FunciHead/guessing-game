package game

import (
	"fmt"
	"guessing-game/bubbletea"
	"guessing-game/core"
	"guessing-game/utils"
	"os"
)

// anti magic values consts

const ATTEMPTS = 5
const GUESS = 0
const LASTGUESS = 0
const SELECTEDGUESS = -1

func Loop() {
	for {
		var myBreak core.LoopBreaker = false
		utils.Cleaner()
		name := utils.DecideName()
		m := bubbletea.Model{Guess: 0, Attempts: ATTEMPTS, SecretNumber: core.MakeRandomNumber(), LastGuess: LASTGUESS, Hint: '?', Name: name, SelectedGuess: SELECTEDGUESS, IsWin: false, Breaker: &myBreak}
		p := bubbletea.NewProgramExport(m)

		if _, err := p.Run(); err != nil {
			fmt.Printf("There's been an error: %v", err)
			os.Exit(1)
		}
		if myBreak {
			utils.Cleaner()
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
