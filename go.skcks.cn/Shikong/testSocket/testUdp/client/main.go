package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8848,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = socket.Close()
	}()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		inputData, _ := inputReader.ReadString('\n')
		inputString := strings.Trim(inputData, "\r\n")

		sendData := []byte(inputString)
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := make([]byte, 4096)
		l, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:l]), remoteAddr)
	}
}
