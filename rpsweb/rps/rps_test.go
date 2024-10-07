package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)
		switch i {
		case 0:
			fmt.Println("El jugador eligio ROCK")
		case 1:
			fmt.Println("El jugador eligio PAPER")
		case 2:
			fmt.Println("El jugador eligio SCISSORS")
		}
		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("Computer choice: %s\n", round.ComputerChoice)
		fmt.Printf("Round result: %s\n", round.RoundResult)
		fmt.Printf("Computer choice int: %d\n", round.ComputerChoiceInt)
		fmt.Printf("Computer score: %d\n", round.ComputerScore)
		fmt.Printf("Player score: %d\n", round.PlayerScore)
	}
}
