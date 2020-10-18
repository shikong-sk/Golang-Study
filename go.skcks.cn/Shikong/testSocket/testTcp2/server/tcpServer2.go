package main

import (
	"bufio"
	"fmt"
	"net"
	proto "testSocket/customProto"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()
	for {
		reader := bufio.NewReader(conn)

		recvStr, err := proto.Decode(reader)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("来自客户端的数据 =>", recvStr)
		_, _ = conn.Write([]byte(recvStr))
	}
}