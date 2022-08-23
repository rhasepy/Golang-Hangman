package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func UtilMsg() string {
	return "Hello, I'm Util Module..."
}

func ReadInput_Int() int {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if one_str := "1\n"; strings.Compare(one_str, text) == 0 {
		return 1
	} else if two_str := "2\n"; strings.Compare(two_str, text) == 0 {
		return 2
	} else {
		return 0
	}
}

func ReadInput_Char() byte {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return byte(text[0])
}

func ReadFileContent(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("[ERROR] File Open: %s\n", path)
		os.Exit(0)
	}
	defer file.Close()

	var WordList []string
	for reader := bufio.NewReader(file); ; {

		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		WordList = append(WordList, strings.ToLower(string(line)))
	}

	return WordList
}

func ContainsBytes(arr []byte, target byte) bool {

	if arr == nil {
		return false
	}

	if len(arr) == 0 {
		return false
	}

	for _, item := range arr {
		if byte(item) == target {
			return true
		}
	}

	return false
}

func GetCurrentTime() string {
	return time.Now().Format("01-02-2006 15:04:05")
}
