package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8848,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = listen.Close()
	}()
	for {
		data := new([1024]byte)
		l, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(data[:l]), addr)
		_, err = listen.WriteToUDP(data[:l], addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
