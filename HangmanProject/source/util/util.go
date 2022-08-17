package util

import (
	"bufio"
	"os"
	"strings"
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
