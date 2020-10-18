package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	proto "testSocket/customProto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")

		// 使用自定义协议头解决 TCP 黏包
		data, err := proto.Encode(inputInfo)

		if err != nil {
			fmt.Println(err)
			return
		}

		for i := 0; i < 10; i++ {
			_, err = conn.Write(data)
		}

		buf := [1024]byte{}
		l, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:l]))
	}
}
