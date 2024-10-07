package rps

import (
	"math/rand"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     int    `json:"computer_score"`
	PlayerScore       int    `json:"player_score"`
}

var winMessages = []string{
	"Bien Hecho",
	"Buen Trabajo",
	"Deberías comprar un boleto de lotería",
}

var loseMessages = []string{
	"Que lastima",
	"Intentalo de nuevo",
	"Hoy simplemente no es tu día",
}

var drawMessages = []string{
	"Empate",
	"Intentar de nuevo",
	"¡Nadie gana, pero puedes intentarlo de nuevo!",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoice = "la pc eligió ROCK"
		computerChoiceInt = ROCK
		if playerValue == ROCK {
			roundResult = "Tie"
		} else if playerValue == PAPER {
			roundResult = "Loser"
		} else {
			roundResult = "Winner"
		}
	case PAPER:
		computerChoice = "la pc eligió PAPER"
		computerChoiceInt = PAPER
		if playerValue == ROCK {
			roundResult = "Loser"
		} else if playerValue == PAPER {
			roundResult = "Tie"
		} else {
			roundResult = "Winner"
		}
	case SCISSORS:
		computerChoice = "la pc eligió SCISSORS"
		computerChoiceInt = SCISSORS
		if playerValue == ROCK {
			roundResult = "Winner"
		} else if playerValue == PAPER {
			roundResult = "Loser"
		} else {
			roundResult = "Tie"
		}
	}

	messageInt := rand.Intn(3)
	var message string
	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "El computador gana"
		message = loseMessages[messageInt]
	}
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     ComputerScore,
		PlayerScore:       PlayerScore,
	}
}
