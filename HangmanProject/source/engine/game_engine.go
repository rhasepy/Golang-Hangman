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

func gameMenuMsg() {

	fmt.Printf("\nYour Health: %d\n", 5)
	fmt.Print("Choice: ")
}

func EngineRoutine() {

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

	fmt.Println(repository.RepoMsg())
	for {

		gameMenuMsg()

		choice := util.ReadInput_Char()

		fmt.Printf("Your choice: %c\n", choice)
	}
}
