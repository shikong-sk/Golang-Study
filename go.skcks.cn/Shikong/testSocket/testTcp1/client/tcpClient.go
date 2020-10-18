package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

		// 短时间内发送多个数据包会导致 TCP 黏包
		for i := 0; i < 10; i++ {
			_, err := conn.Write([]byte(inputInfo))
			if err != nil {
				fmt.Println(err)
				return
			}
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
