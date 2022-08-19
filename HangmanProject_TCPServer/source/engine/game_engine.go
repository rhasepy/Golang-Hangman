package engine

import (
	"fmt"
	"game/source/Network"
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
	Client       Network.Socket
}

func CreateGameObj(word string, globalScore int, client Network.Socket) GameObj {
	return GameObj{
		WrongMoveCtr: 0,
		Word:         word,
		IsWin:        false,
		Score:        globalScore,
		Client:       client,
	}
}

func ReadMove() byte {

	fmt.Printf("Input: ")
	choice := util.ReadInput_Char()
	return choice
}

func (g *GameObj) DrawHang(totalHeath int) {

	g.Client.WriteSock("I---\n")
	g.Client.WriteSock("   |\n")
	g.Client.WriteSock("  ---\n")

	if g.WrongMoveCtr > 0 {
		g.Client.WriteSock("   O\n")
	}

	for i := 0; i < g.WrongMoveCtr; i++ {

		if i == 1 {
			g.Client.WriteSock("  /|\\\n")
		} else if i == totalHeath-1 {
			g.Client.WriteSock("  / \\\n")
		} else if i < totalHeath && i > 1 {
			g.Client.WriteSock("   |\n")
		}
	}
}

func (g *GameObj) Drawing() {

	g.Client.WriteSock("\n*************************************\n")
	g.DrawHang(config.Health)

	g.Client.WriteSock("\n_____________________________________\n")
	g.Client.WriteSock("Your Word: ")

	isWinFlag := true
	for _, char := range g.Word {

		if util.ContainsBytes(g.TrueMoves, byte(char)) {
			g.Client.WriteSock(fmt.Sprintf("%c", char))
		} else {
			g.Client.WriteSock("_ ")
			isWinFlag = false
		}
	}
	g.Client.WriteSock("\n")
	g.Client.WriteSock(fmt.Sprintf("Your Health: %d\n", config.Health-g.WrongMoveCtr))
	g.Client.WriteSock(fmt.Sprintf("Your Total Score: %d\n", g.Score))
	g.Client.WriteSock("_____________________________________\n")
	g.Client.WriteSock("*************************************\n")

	g.IsWin = isWinFlag
}

func (g *GameObj) CheckMove(move byte) bool {

	if strings.Contains(g.Word, string(move)) {

		if !util.ContainsBytes(g.TrueMoves, move) {
			g.TrueMoves = append(g.TrueMoves, move)
		}
		return true
	}

	return false
}

func GameRoutine(client Network.Socket) {

	repository.Construct()

	stats := UserStats{
		Score: 0,
	}

	for i := 0; ; i++ {

		idx, msg := repository.GetOneWord()
		if idx == -1 {
			fmt.Printf("\n[%s] %s\n", util.GetCurrentTime(), msg)
			os.Exit(0)
		}
		client.WriteSock("\n\n")

		gameObj := CreateGameObj(msg, stats.Score, client)

		for {
			gameObj.Drawing()
			if gameObj.IsWin {
				stats.Score += 1
				gameObj.Score = stats.Score
				gameObj.Client.WriteSock("\nYou win, for this word!\n")
				gameObj.Client.WriteSock("Other word is loading...")
				break
			}

			gameObj.Client.WriteSock("Input: ")
			move := gameObj.Client.ReadSock()[0]

			if !gameObj.CheckMove(move) {
				gameObj.WrongMoveCtr += 1
			}

			if gameObj.WrongMoveCtr > config.Health {
				gameObj.Client.WriteSock("\nYou lose, for this word!\n")
				gameObj.Client.WriteSock("Other word is loading...")
				break
			}
		}
		repository.ScoreWord(idx)

		if i >= repository.GetWordCount() {

			gameObj.Client.WriteSock(fmt.Sprintf("\n\nThe Game Finished...\nYour Score: %d/%d\n",
				stats.Score, repository.GetWordCount()))
			break
		}
	}
}
