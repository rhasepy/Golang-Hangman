package repository

import (
	"game/source/config"
	"game/source/util"
	"math/rand"
	"time"
)

var wordList []string
var scoredWordCount int = 0
var global_idx int = -1
var scoreTable map[int]bool

func Construct() {

	rand.Seed(time.Now().Unix())
	scoreTable = make(map[int]bool)
	wordList = util.ReadFileContent(config.FilePath)

	rand.Shuffle(len(wordList), func(i, j int) {
		wordList[i], wordList[j] = wordList[j], wordList[i]
	})

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

	global_idx += 1
	return global_idx, wordList[global_idx]
}
