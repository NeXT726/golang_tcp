package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type conData struct {
	Addr      uint32
	RemoteKey uint32
}

func main() {

	var clientConData = conData{0, 000}

	tcpServer, _ := net.ResolveTCPAddr("tcp", ":19875")

	listener, _ := net.ListenTCP("tcp", tcpServer)

	for {
		conn, _ := listener.Accept()
		handle(conn, &clientConData)
		fmt.Println(clientConData)
	}
}

func handle(conn net.Conn, cD *conData) {
	defer conn.Close()

	data := make([]byte, 8)
	reader := bufio.NewReader(conn)
	reader.Read(data)
	buf := bytes.NewBuffer(data)
	err := binary.Read(buf, binary.LittleEndian, cD)
	if err != nil {
		fmt.Println("wrong:", err)
	}

	time.Sleep(1 * time.Second)
	now := time.Now().String()
	conn.Write([]byte(now))
}
