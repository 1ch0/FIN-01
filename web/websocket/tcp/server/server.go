package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	checkError(err)
	listener, err := net.Listen("tcp4", tcpAddr.String())
	checkError(err)
	conn, err := listener.Accept()
	checkError(err)

	request := make([]byte, 256)
	n, err := conn.Read(request)
	checkError(err)

	fmt.Printf("request is: %s \n", string(request[:n]))

	reponse := "hello" + string(request[:n])
	_, err = conn.Write([]byte(reponse))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
