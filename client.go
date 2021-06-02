package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type conData struct {
	Addr      uint32
	RemoteKey uint32
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Wrong input!")
		return
	}
	server := os.Args[1]
	addr, _ := net.ResolveTCPAddr("tcp", server+":19875")
	conn, _ := net.DialTCP("tcp", nil, addr)

	var clientConData = conData{11, 111}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, clientConData)
	if err != nil {
		fmt.Println("write to buffer wrong:", err)
	}
	_, _ = conn.Write(buf.Bytes())

	response, _ := ioutil.ReadAll(conn)
	fmt.Println(string(response))
}
