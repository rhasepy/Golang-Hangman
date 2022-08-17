package repository

import (
	"game/source/util"
	"math/rand"
)

var _FilePath string = "resource/wordlist.txt"

var wordList []string
var takenWordCount int = 0
var scoredWordCount int = 0
var scoreTable map[int]bool

func Construct() {

	scoreTable = make(map[int]bool)
	wordList = util.ReadFileContent(_FilePath)
	for index := range wordList {
		scoreTable[index] = false
	}
}

func ScoreWord(wordID int) {

	if _, ok := scoreTable[wordID]; ok {
		scoreTable[wordID] = true
	}

	scoredWordCount += 1
}

func GetOneWord() (int, string) {

	if takenWordCount == len(scoreTable) {

		if scoredWordCount == len(scoreTable) {
			return -1, "You Win, Congratulations!"
		} else {
			return -1, "All word are gone, You Lose!"
		}

	} else if len(wordList) == 0 {
		return -1, "[ERROR] Wordlist is empty!"
	}

	for {
		idx := rand.Intn(len(wordList))
		if !scoreTable[idx] {
			takenWordCount += 1
			return idx, wordList[idx]
		}
	}
}
