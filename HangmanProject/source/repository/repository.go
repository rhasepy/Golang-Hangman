package repository

import (
	"game/source/config"
	"game/source/util"
	"math/rand"
	"time"
)

var wordList []string
var takenWordCount int = 0
var scoredWordCount int = 0
var scoreTable map[int]bool

func Construct() {

	scoreTable = make(map[int]bool)
	wordList = util.ReadFileContent(config.FilePath)
	for index := range wordList {
		scoreTable[index] = false
	}
}

func GetWordCount() int {
	return len(wordList)
}

func ScoreWord(wordID int) {

	if _, ok := scoreTable[wordID]; ok {
		scoreTable[wordID] = true
	}

	scoredWordCount += 1
}

func GetOneWord() (int, string) {

	if len(wordList) == 0 {
		return -1, "[ERROR] Wordlist is empty!"
	}

	for {
		rand.Seed(time.Now().Unix())
		idx := rand.Intn(len(wordList))
		if !scoreTable[idx] {
			takenWordCount += 1
			return idx, wordList[idx]
		}
	}
}
