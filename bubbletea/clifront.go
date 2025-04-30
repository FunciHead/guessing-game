package bubbletea

import (
	"fmt"
	"guessing-game/core"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Guess         int
	Attempts      int
	SecretNumber  int
	LastGuess     int
	Name          string
	Hint          rune
	SelectedGuess int
	IsWin         bool
	Breaker       *core.LoopBreaker
}

func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) tryGuess(guess int, SecretNumber int) bool {

	return core.CheckGuess(guess, SecretNumber)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.Guess < core.RandNUm {
				m.Guess++
			}
		case "down":
			if m.Guess > 0 {
				m.Guess--
			}
		case "enter":
			if m.IsWin || m.Attempts <= 0 {
				return m, tea.Quit

			} else {
				m.guessProcessing(&m)
			}
		case "crtl+c", "q":
			*m.Breaker = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) guessProcessing(model *Model) {

	m.SelectedGuess = m.Guess
	if m.tryGuess(m.SelectedGuess, m.SecretNumber) {
		model.IsWin = true
	} else {
		model.Hint = core.CheckSide(m.Guess, m.SecretNumber)
		model.LastGuess = m.Guess
		model.Attempts--
	}
}

func (m Model) View() string {

	if m.IsWin {
		return fmt.Sprintf("You won!!! %d was the secret number\nPress \"enter\" to play again\nPress\"q\" to quit ", m.SecretNumber)

	} else if !m.IsWin && m.Attempts > 0 {
		return fmt.Sprintf("\n\"↑/↓\" to change guess, \"q\" to quit, \"enter\" to guess \n%s you have Attempts: %d guess: %d\n(HINT) Last guess: %d %c secret number  ", m.Name, m.Attempts, m.Guess, m.LastGuess, m.Hint)

	} else {
		return fmt.Sprintf("You lost!!! %d was the secret number\nPress \"enter\" to play again\nPress\"q\" to quit ", m.SecretNumber)
	}

}

func NewProgramExport(m Model) *tea.Program {
	return tea.NewProgram(m)
}
