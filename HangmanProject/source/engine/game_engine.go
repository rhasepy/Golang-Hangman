package engine

import (
	"fmt"
	"game/source/repository"
	"game/source/util"
	"os"
)

func mainMenuMsg() {

	fmt.Printf("\n\n")
	fmt.Println("*** *** MENU *** ***")
	fmt.Println("1) Start Hangman Game")
	fmt.Println("2) Exit")

	fmt.Print("Choice: ")
}

func StartGameActivity() {

	fmt.Printf("Game Started...\n")

	for {

		mainMenuMsg()

		choice := util.ReadInput_Int()

		switch choice {
		case 1:
			GameRoutine()
		case 2:
			fmt.Println("Good Bye!")
			os.Exit(0)
		default:
			fmt.Println("Please enter the valid input (1 or 2)...")
		}
	}
}

func GameRoutine() {

	repository.Construct()

	for gameLoop := true; gameLoop; {

		idx, msg := repository.GetOneWord()
		if idx == -1 {
			gameLoop = false
			fmt.Printf("\n%s\n", msg)
			os.Exit(0)
		}

		/*gameMenuMsg()*/

		fmt.Printf("Input: ")
		choice := util.ReadInput_Char()
		fmt.Printf("%c, Processing...\n", choice)
	}
}
