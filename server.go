package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"encoding/json"
)

var count = 0
var secretNumber = 0

func handleConnection(c net.Conn) {
	for {
		netData, connErr := bufio.NewReader(c).ReadString('\n')
		if connErr != nil {
			fmt.Println(connErr)
			return
		}

		temp := strings.ToUpper(strings.TrimSpace(string(netData)))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func gameStart(message string) {
	message := json {
		message:	message,
		timestamp:	time.Now().Unix(),
		gameId: 1
	}
}
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	listener, listenerr := net.Listen("tcp4", PORT)
	if listenerr != nil {
		fmt.Println(listenerr)
		return
	}
	defer listener.Close()

	//Generate the secret number
	rand.Seed(time.Now().UnixNano())
	minSecretNum := 0
	maxSecretNum := 500
	secretNumber = rand.Intn(maxSecretNum-minSecretNum+1) + minSecretNum
	fmt.Println(secretNumber)
	// fmt.Println("secretNumber: " + secretNumber)

	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}

	// for {
	// 	netData, err := bufio.NewReader(c).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if strings.TrimSpace(string(netData)) == "STOP" {
	// 		fmt.Println("Exiting TCP server!")
	// 		return
	// 	}

	// 	fmt.Print("-> ", string(netData))
	// 	t := time.Now()
	// 	myTime := t.Format(time.RFC3339) + "\n"
	// 	c.Write([]byte(myTime))
	// }
}
