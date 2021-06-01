package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	tcpServer, _ := net.ResolveTCPAddr("tcp", "10.251.31.97:19875")

	listener, _ := net.ListenTCP("tcp", tcpServer)

	for {
		conn, _ := listener.Accept()
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	go func() {
		response, _ := ioutil.ReadAll(conn)
		fmt.Println(string(response))
	}()

	time.Sleep(1 * time.Second)
	now := time.Now().String()
	conn.Write([]byte(now))
}
