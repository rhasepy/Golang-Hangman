package engine

import (
	"fmt"
	"game/source/config"
	"game/source/repository"
	"game/source/util"
	"os"
	"strings"
)

type UserStats struct {
	Score int
}

type GameObj struct {
	WrongMoveCtr int
	Word         string
	TrueMoves    []byte
	IsWin        bool
	Score        int
	IsStarted    bool
}

func CreateGameObj(word string, globalScore int) GameObj {
	return GameObj{
		WrongMoveCtr: 0,
		Word:         word,
		IsWin:        false,
		Score:        globalScore,
		IsStarted:    true,
	}
}

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

func ReadMove() byte {

	fmt.Printf("Input: ")
	choice := util.ReadInput_Char()
	return choice
}

func (g *GameObj) DrawHang(totalHeath int) {

	fmt.Printf("I---\n")
	fmt.Printf("   |\n")
	fmt.Printf("  ---\n")

	if !g.IsStarted {

		if g.WrongMoveCtr > 0 {
			fmt.Printf("   O\n")
		}

		for i := 0; i < g.WrongMoveCtr; i++ {

			if i == 1 {
				fmt.Printf("  /|\\\n")
			} else if i == totalHeath-1 {
				fmt.Printf("  / \\\n")
			} else if i < totalHeath && i > 1 {
				fmt.Printf("   |\n")
			}
		}

	}
}

func (g *GameObj) Drawing() {

	fmt.Printf("\n*************************************\n")
	g.DrawHang(config.Health)

	fmt.Printf("\n_____________________________________\n")
	fmt.Printf("Your Word: ")

	isWinFlag := true
	for _, char := range g.Word {

		if util.ContainsBytes(g.TrueMoves, byte(char)) {
			fmt.Printf("%c", char)
		} else {
			fmt.Printf("_ ")
			isWinFlag = false
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Your Health: %d\n", config.Health-g.WrongMoveCtr)
	fmt.Printf("Your Total Score: %d\n", g.Score)
	fmt.Printf("_____________________________________\n")
	fmt.Printf("*************************************\n")

	g.IsWin = isWinFlag
}

func (g *GameObj) CheckMove(move byte) bool {

	if strings.Contains(g.Word, string(move)) {

		if !util.ContainsBytes(g.TrueMoves, move) {
			g.TrueMoves = append(g.TrueMoves, move)
		}
		return true
	}

	g.IsStarted = false
	return false
}

func GameRoutine() {

	repository.Construct()

	stats := UserStats{
		Score: 0,
	}

	for gameLoop := true; gameLoop; {

		idx, msg := repository.GetOneWord()
		if idx == -1 {
			gameLoop = false
			fmt.Printf("\n%s\n", msg)
			os.Exit(0)
		}
		fmt.Printf("\n\n")

		game_obj := CreateGameObj(msg, stats.Score)

		for {

			game_obj.Drawing()

			if game_obj.IsWin {
				game_obj.Score += 1
				fmt.Printf("\nYou win, for this word!\n")
				fmt.Printf("Other word is loading...")
				break
			}

			move := ReadMove()

			if !game_obj.CheckMove(move) {
				game_obj.WrongMoveCtr += 1
			}

			if game_obj.WrongMoveCtr > config.Health {
				fmt.Printf("\nYou lose, for this word!\n")
				fmt.Printf("Other word is loading...")
				break
			}
		}

		repository.ScoreWord(idx)
	}

	fmt.Printf("The Game Finished...\nYour Score: %d/%d\n",
		stats.Score, repository.GetWordCount())
}
