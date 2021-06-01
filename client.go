package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong input!")
		return
	}

	server := os.Args[1]
	addr, _ := net.ResolveTCPAddr("tcp", server+":19875")
	conn, _ := net.DialTCP("tcp", nil, addr)
	_, _ = conn.Write([]byte("Hello TCP\n"))
	response, _ := ioutil.ReadAll(conn)
	fmt.Println(string(response))
}
