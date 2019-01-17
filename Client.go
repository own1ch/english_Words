package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

//Client.go for tests

const (
	message       = "PING"
	StopCharacter = "STOP"
)

type result struct {
	id            int
	word          string
	transcription string
	translation   string
}

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

	buff := make([]byte, 1024)
	for {
		reader := bufio.NewReader(os.Stdin)
		sms, _ := reader.ReadString('\n')
		if strings.TrimRight(sms, "\n") == "stop client" {
			break
		}
		con.Write([]byte(sms))
		con.Read(buff)
		bin_buf := bytes.NewBuffer(buff)
		var answer = new(result)
		obj := gob.NewDecoder(bin_buf)
		obj.Decode(answer)
		log.Println(answer.word)
	}
}
