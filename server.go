package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":7")
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	HandleRequest(buffer[:n], conn, err)
}

func HandleRequest(data []byte, conn net.Conn, err error) {

	var decoded = interface{}(nil)
	err = msgpack.Unmarshal(data, &decoded)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}

	fmt.Println("Received data:", decoded)

	response := []byte(fmt.Sprint(decoded))
	_, err = conn.Write(response)
	if err != nil {
		fmt.Println("Error sending response:", err)
		return
	}
}
