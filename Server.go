package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"./ServerCommands"
)

func main() {
	/*arguments := os.Args*/
	/*if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}*/
	PORT := ":8080"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, errorRead := bufio.NewReader(c).ReadString('\n')
		if errorRead != nil {
			fmt.Println(errorRead)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			os.Exit(3)
		}

		if temp == ServerCommands.GET_WORDS {
			ServerCommands.Commands(ServerCommands.GET_WORDS)
		}

		result := strconv.Itoa(5) + "\n"
		_, errorWrite := c.Write([]byte(string(result)))
		if errorWrite != nil {
			fmt.Println(errorWrite)
			return
		}

		fmt.Println(netData)
	}
	defer c.Close()

	fmt.Println("server closed")
}
