package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type clientRequest struct {
	message string
	playerName string
	timestamp int64
}

type serverResponse struct {
	message string
	playerName string
	timestamp int64
	gameId int64
}

type guessMessage struct {
	message string
	guess int
	timestamp int64
	gameId int64
}

type guessResult struct {
	message string
	guessResult int
	timestamp int64
	gameId int64
}

type winMessage struct {
	message string
	answer int
	winner string
	gameId int64
}

type gameStartMessage struct {
	message string
	timestamp int64
	gameId int64
}

type errorMessage struct {
	message string
	reason string
	timestamp int64
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Host:port is missing.")
		return
	}

	connect := arguments[1]
	conn, connErr := net.Dial("tcp", connect)
	if connErr != nil {
		fmt.Println(connErr)
		return
	} else {
		fmt.Println(conn.)
	}

	for {		
		//guesting answer
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter the number :\n")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message: " + message)
		text = strings.ToUpper(strings.TrimSpace(string(text)))
		if text == "STOP" {
			fmt.Println("TCP client exit")
			return
		}

		fmt.Println(text)

		intNum, intErr := strconv.Atoi(text)
		if intErr == nil {
			if intNum == 123 {
				fmt.Println("You won, the number is 123")
				return
			} else {
				fmt.Println("Text: " + text + ", intNum: " + string(intNum))
				// return
			}
		} else {
			fmt.Println("Text: " + text + " is not an integer." + fmt.Sprint(intErr))
			// fmt.Println(intErr)
		}

	}

}
