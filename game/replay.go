package game

import (
	"bufio"
	"fmt"
	"os"
)

func PlayAgain() bool {
	for {
		fmt.Println("Do you want to play again?\n[Y]es\n[N]o")
		var want rune
		_, err := fmt.Scanf("%c", &want)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("INPUT ERROR, try again!")
			continue
		} else {
			bufio.NewReader(os.Stdin).ReadString('\n')
			if want == 'n' {
				return false
			} else if want == 'y' {
				return true
			} else {
				continue
			}
		}
	}
}
