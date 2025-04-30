package main

import (
	"fmt"
	"guessing-game/game"
	"guessing-game/utils"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var finishLoopGlobalVariable bool

type model struct {
	guess         int
	attempts      int
	secretNumber  int
	lastGuess     int
	name          string
	hint          rune
	selectedGuess int
	isWin         bool
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) tryGuess(guess int, secretNumber int) bool {

	return game.CheckGuess(guess, secretNumber)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.guess < game.RandNUm {
				m.guess++
			}
		case "down":
			if m.guess > 0 {
				m.guess--
			}
		case "enter":
			if !m.isWin {
				m.selectedGuess = m.guess
				if m.tryGuess(m.selectedGuess, m.secretNumber) {
					m.isWin = true
				} else {
					m.hint = game.CheckSide(m.guess, m.secretNumber)
					m.lastGuess = m.guess
					m.attempts--

				}

			} else {
				return m, tea.Quit
			}

		case "crtl+c", "q":
			finishLoopGlobalVariable = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {

	if !m.isWin {
		return fmt.Sprintf("\n\"↑/↓\" to change guess, \"q\" to quit, \"enter\" to guess \n%s you have Attempts: %d guess: %d\n(HINT) Last guess: %d %c secret number  ", m.name, m.attempts, m.guess, m.lastGuess, m.hint)
	} else {
		return fmt.Sprintf("You won!!! %d was the secret number\nPress \"enter\" to play again\nPress\"q\" to quit ", m.secretNumber)
	}

}

func main() {
	for {
		utils.Cleaner()
		finishLoopGlobalVariable = false
		name := game.DecideName()
		m := model{guess: 0, attempts: 1, secretNumber: game.MakeRandomNumber(), lastGuess: 0, hint: '?', name: name, selectedGuess: -1, isWin: false}
		p := tea.NewProgram(m)

		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
		if finishLoopGlobalVariable {
			println("Thank you very much")
			break
		}
	}
	utils.Cleaner()
}
