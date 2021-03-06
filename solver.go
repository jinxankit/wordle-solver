package main

import (
	"fmt"
	"strings"
)

type solver struct {
	game *Game
}

func (s *solver) makeAGuess() string {
	var guess string
	guessesLeft := MaxGuessCount - s.game.guessCount
	fmt.Printf("%d guesses left. Please make a guess...... ", guessesLeft)
	fmt.Scanf("%s", &guess)
	guess = strings.ToLower(strings.TrimSpace(guess))
	return guess
}

func (s *solver) onEndOfGame() {
	if s.game.wonStatus {
		fmt.Println("You won in ", s.game.guessCount, " guesses!")
		fmt.Println("The word was: ", s.game.targetWord)
	} else {
		fmt.Println("You are out of guesses. The correct word was:", s.game.targetWord)
	}
}
