package main

import (
	"bufio"
	"fmt"
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
		data := make([]byte, 10)
		reader := bufio.NewReader(conn)
		reader.Read(data)

		//data, _ := ioutil.ReadAll(conn)
		fmt.Println(string(data), "\t", len(data), "\t", cap(data))
	}()

	time.Sleep(1 * time.Second)
	now := time.Now().String()
	conn.Write([]byte(now))
}
