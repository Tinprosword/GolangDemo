package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type GamePlay struct {
	PlayerId   int
	PlayerName string
	Message    string
	Timestamp  int64
	GameId     int
	Answer     uint
}

// Package variables for all files
var SecretAnswer uint
var MinAnswerRange uint = 1
var MaxAnswerRange uint = 500
var Players = make([]GamePlay, 0)

func main() {
	// Generate the random secret answer
	var answerRange = int(MaxAnswerRange - MinAnswerRange)
	rand.Seed(time.Now().UTC().UnixNano())
	var gameId = rand.Int()
	SecretAnswer = uint(rand.Intn(answerRange)) + MinAnswerRange
	fmt.Printf("The secret answer of Game Id %v is %v\n", gameId, SecretAnswer)

	for {
		var playerName string
		var answer uint
		var userInput int
		var playerId = rand.Int()
		var playerIndex int
		var gameTimestamp = time.Now().UTC().UnixNano()
		var player = GamePlay{
			PlayerId:  playerId,
			Timestamp: gameTimestamp,
			GameId:    gameId,
		}
		fmt.Println("Welcome to the game, please enter your name to start.")
		// fmt.Scan does not support spaces so changed to bufio.Reader and os.Stdin
		// fmt.Scan(&playerName)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		playerName = scanner.Text()

		if !validatePlayer(playerName) {
			fmt.Println("The name is invalid, please try again.")
			continue
		}
		player.PlayerName = playerName
		Players = append(Players, player)
		playerIndex = len(Players) - 1

		fmt.Printf("There are %v players, your index is %v\n", len(Players), playerIndex)

		fmt.Printf("Welcome %v, please guess the number between %v and %v.\n", playerName, MinAnswerRange, MaxAnswerRange)
		for {
			// wait user to input answer
			fmt.Scan(&userInput)
			// scanner.Scan()
			// t := scanner.Text()
			// answer = uint(t)
			var isValidAnswer = validateAnswer(userInput)
			if !isValidAnswer {
				fmt.Println("The answer is invalid, please try again.")
				continue
			} else {
				answer = uint(userInput)
			}
			// fmt.Println("=======================")
			player.Answer = answer
			Players[playerIndex].Answer = answer

			// player = verifyAnswer(player)
			// passing by pointer
			verifyAnswer(&player)
			fmt.Println(player.Message)
			// fmt.Printf("Secret is %v and guess is %v\n", SecretAnswer, player.Answer)
			Players[playerIndex].Message = player.Message

			// Game ending condition
			if strings.Contains(player.Message, "Congratulation") {
				for i, _ := range Players {
					Players[i].Message = player.Message
				}
				break
			} else {
				fmt.Printf("Please guess the number between %v and %v.\n", MinAnswerRange, MaxAnswerRange)
			}
		}
		fmt.Printf("All players: %v\n", Players)

		// End game
		break
		// New game logic here
	}
}

func validatePlayer(player string) bool {
	if len(player) < 2 {
		return false
	}
	return true
}

func validateAnswer(answer int) bool {
	if answer < int(MinAnswerRange) {
		return false
	} else if answer > int(MaxAnswerRange) {
		return false
	}
	return true
}

func findPlayer(player GamePlay) int {
	// return index in Players
	for i, p := range Players {
		if p.PlayerId == player.PlayerId {
			return i
		}
	}
	return -1
}

func verifyAnswer(gamePlay *GamePlay) *GamePlay {
	// func verifyAnswer(gamePlay GamePlay) GamePlay {
	var message string
	var answer = gamePlay.Answer
	var player = gamePlay.PlayerName
	if answer == SecretAnswer {
		message = fmt.Sprintf("Congratulation to %v, the answer is %v\n", player, answer)
	} else if answer > SecretAnswer {
		message = "Too large"
	} else {
		message = "Too small"
	}
	gamePlay.Message = message
	return gamePlay
}
