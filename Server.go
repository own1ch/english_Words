package main

import (
	"./ServerCommands"
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
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
			result := ServerCommands.Commands(ServerCommands.GET_WORDS)
			binBuf := new(bytes.Buffer)
			obj := gob.NewEncoder(binBuf)
			obj.Encode(result)
			_, errorWrite := c.Write(binBuf.Bytes())
			if errorWrite != nil {
				fmt.Println(errorWrite)
				return
			}

			fmt.Println(netData)
			fmt.Println(result)
		}
	}
	defer c.Close()

	fmt.Println("server closed")
}
