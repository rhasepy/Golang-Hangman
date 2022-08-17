package main

import (
	"fmt"
	"game/source/engine"
)

func main() {
	fmt.Println("******* ****** ****** ****** ****** ****** ******")
	fmt.Println("*** *** *** Welcome the Hangman Game! *** *** ***")
	fmt.Println("*** *** *** * * Game Starting * * *** *** *** ***")
	fmt.Println("******* ****** ****** ****** ****** ****** ******")
	fmt.Printf("\n\n")

	engine.StartGameActivity()
}
