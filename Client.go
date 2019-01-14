package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

//Client.go for tests

const (
	message = "PING"
	StopCharacter = "STOP"
)

func main() {
	var ip string = "127.0.0.1"
	var port int = 8080

	Client(ip, port)
}

func Client(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	con, err := net.Dial("tcp", addr)
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		sms, _ := reader.ReadString('\n')
		if strings.TrimRight(sms, "\n") == "stop client" {
			break
		}
		con.Write([]byte(sms))
		log.Printf("sent %s", message)

		buff := make([]byte, 1024)
		n, _ := con.Read(buff)
		log.Printf("Receive %s", n)
	}
}
