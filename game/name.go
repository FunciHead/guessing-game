package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DecideName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your name: ")
	for {
		name, err := reader.ReadString('\n')
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("INPUT ERROR: ", err)
			fmt.Println("Try again!")
			continue
		} else {
			name = strings.TrimSpace(name)

			return name
		}
	}
}
