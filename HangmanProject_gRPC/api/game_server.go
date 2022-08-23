package api

import (
	"fmt"
	"hangman_grpc/source/config"
	"hangman_grpc/source/repository"
	"hangman_grpc/util"
	"os"
	"strings"
	"sync"
)

type gameMessage struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type gameHandle struct {
	MQue  []gameMessage
	mutex sync.Mutex
}

var gameHandleObject = gameHandle{}

type GameServer struct {
}

func (is *GameServer) mustEmbedUnimplementedServicesServer() {
	panic("unsupported exception")
}

// This Context run client specific
// Actually this method handshake with specific client on specific context
func (is *GameServer) GameService(gsi Services_GameServiceServer) error {

	//clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	go handleRequest(gsi, errch)

	// Receive Message
	//go receiveFromStream(gsi, clientUniqueCode, errch)

	// Send Message
	//go sendToStream(gsi, clientUniqueCode, errch)

	return <-errch
}

func handleRequest(client Services_GameServiceServer, errch chan error) {
	GameRoutine(client, errch)
}

type UserStats struct {
	Score int
}

type GameObj struct {
	WrongMoveCtr int
	Word         string
	TrueMoves    []byte
	IsWin        bool
	Score        int
	Client       Services_GameServiceServer
}

func CreateGameObj(word string, globalScore int, client Services_GameServiceServer) GameObj {
	return GameObj{
		WrongMoveCtr: 0,
		Word:         word,
		IsWin:        false,
		Score:        globalScore,
		Client:       client,
	}
}

func (g *GameObj) DrawHang(totalHeath int) {

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "I---\n",
	})

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "   |\n",
	})

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "  ---\n",
	})

	if g.WrongMoveCtr > 0 {
		g.Client.Send(&FromServer{
			Name: "Server",
			Body: "   O\n",
		})
	}

	for i := 0; i < g.WrongMoveCtr; i++ {

		if i == 1 {
			g.Client.Send(&FromServer{
				Name: "Server",
				Body: "  /|\\\n",
			})
		} else if i == totalHeath-1 {
			g.Client.Send(&FromServer{
				Name: "Server",
				Body: "  / \\\n",
			})
		} else if i < totalHeath && i > 1 {
			g.Client.Send(&FromServer{
				Name: "Server",
				Body: "   |\n",
			})
		}
	}
}

func (g *GameObj) Drawing() {

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "\n*************************************\n",
	})
	g.DrawHang(config.Health)

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "\n_____________________________________\n",
	})
	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "Your Word: ",
	})

	isWinFlag := true
	for _, char := range g.Word {

		if util.ContainsBytes(g.TrueMoves, byte(char)) {
			g.Client.Send(&FromServer{
				Name: "Server",
				Body: fmt.Sprintf("%c", char),
			})
		} else {
			g.Client.Send(&FromServer{
				Name: "Server",
				Body: "_ ",
			})
			isWinFlag = false
		}
	}

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "\n",
	})

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: fmt.Sprintf("Your Health: %d\n", config.Health-g.WrongMoveCtr),
	})

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: fmt.Sprintf("Your Total Score: %d\n", g.Score),
	})

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "_____________________________________\n",
	})

	g.Client.Send(&FromServer{
		Name: "Server",
		Body: "*************************************\n",
	})

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

func sendData(server Services_GameServiceServer, msg string) {

	server.Send(&FromServer{
		Name: "Server",
		Body: msg,
	})
}

func GameRoutine(client Services_GameServiceServer, errch chan error) error {

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
		sendData(client, "\n\n")

		gameObj := CreateGameObj(msg, stats.Score, client)

		for {
			gameObj.Drawing()
			if gameObj.IsWin {
				stats.Score += 1
				gameObj.Score = stats.Score

				sendData(gameObj.Client, "\nYou win, for this word!\n")
				sendData(gameObj.Client, "Other word is loading...")
				break
			}

			sendData(gameObj.Client, "Input: ")
			msg, err := gameObj.Client.Recv()
			if err != nil {
				fmt.Fprintf(os.Stdout, "[%s] One client left!\n", util.GetCurrentTime())
				return <-errch
			}

			fmt.Fprintf(os.Stdout, "[%s] Client - %s, Body: %s\n", util.GetCurrentTime(), msg.Name, msg.Body)
			move := msg.Body[0]

			if !gameObj.CheckMove(move) {
				gameObj.WrongMoveCtr += 1
			}

			if gameObj.WrongMoveCtr > config.Health {
				sendData(gameObj.Client, "\nYou lose, for this word!\n")
				sendData(gameObj.Client, "Other word is loading...")
				break
			}
		}
		repository.ScoreWord(idx)

		if i >= repository.GetWordCount() {

			sendData(gameObj.Client, fmt.Sprintf("\n\nThe Game Finished...\nYour Score: %d/%d\n",
				stats.Score, repository.GetWordCount()))
			break
		}
	}

	return <-errch
}
