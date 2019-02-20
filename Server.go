package main

import (
	"./ServerCommands"
	"bufio"
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
	PORT := ":6666"
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

		var command string
		command = string(netData[2 : len(netData)-1])
		if strings.Contains(command, "getWords") {
			data := strings.Split(command, "|")
			countOfWords := strings.TrimSpace(string(data[0][len(data[0])-2:]))
			login := data[1]
			result := ServerCommands.GetWords(countOfWords, login)
			words := strings.Join(result, "|") + "\n"
			_, errorWrite := c.Write([]byte(words))
			if errorWrite != nil {
				fmt.Println(errorWrite)
				return
			}
			fmt.Println(netData)
			fmt.Println(result)
		} else if strings.Contains(command, "registration") {
			/**example netData = "eremin1|a22021980|Виталий|registration" */
			reg := ServerCommands.Registration(netData)
			if reg {
				_, errorWrite := c.Write([]byte("true"))
				CheckErr(errorWrite)
				fmt.Println("Пользователь успешно зарегистрирован!")
			} else {
				_, errorWrite := c.Write([]byte("false"))
				CheckErr(errorWrite)
				fmt.Printf("Пользователь не зарегистрирован! Что-то пошло не так! Возможно пользователь с таким именем уже существует!%s\n", netData)
			}
		} else if strings.Contains(command, "login") {
			/**example netData = "eremin|a22021980|login" **/
			login := ServerCommands.Login(netData)
			if login {
				_, errorWrite := c.Write([]byte("true"))
				CheckErr(errorWrite)
				fmt.Println("Пользователь успешно залогинился!")
			} else {
				_, errorWrite := c.Write([]byte("false"))
				CheckErr(errorWrite)
				fmt.Println("Пользователь не смог залогиниться!")
			}
		} else if command == "STOP" {
			os.Exit(3)
		}

	}
	defer c.Close()

	fmt.Println("server closed")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
