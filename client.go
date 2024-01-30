package main

import (
	"fmt"
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	conn, err := net.Dial("tcp", ":7")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected. Sending data...")

	sendRequest(conn)
}

func sendRequest(conn net.Conn) {
	body, err := msgpack.Marshal(people)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	_, err = conn.Write(body)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}

	fmt.Println("Response:", string(buffer[:n]))
}

type Person struct {
	Name   string
	Age    int
	Height float32
}

var people = []Person{
	{"Akari Akaza", 14, 149.5},
	{"Kyoko Toshino", 14, 154.5},
	{"Yui Funami", 14, 155.0},
	{"Chinatsu Yoshikawa", 14, 150.0},
	{"Sakurako Oomuro", 14, 158.0},
	{"Himawari Furutani", 14, 157.0},
	{"Rise Matsumoto", 14, 155.0},
	{"Nana Nishigaki", 14, 155.0},
	{"Ayano Sugiura", 14, 157.0},

	{"Chino Kafu", 13, 142.0},
	{"Cocoa Hoto", 16, 149.0},
	{"Rize Tedeza", 15, 159.0},
	{"Chiya Ujimatsu", 15, 156.0},
	{"Sharo Kirima", 16, 147.0},
	{"Maya Jouga", 16, 155.0},
	{"Megumi Natsu", 16, 154.0},
	{"Mocha Hoto", 13, 140.0},
	{"Midori Aoyama", 16, 158.0},
	{"Aoyama Blue Mountain", 16, 158.0},
	{"Anko Tedeza", 15, 150.0},
	{"Tippy", 1, 20.0},

	{"Yui Hirasawa", 17, 156.0},
	{"Ritsu Tainaka", 17, 159.0},
	{"Mio Akiyama", 17, 160.0},
	{"Tsumugi Kotobuki", 17, 155.0},
	{"Azusa Nakano", 16, 150.0},
	{"Ui Hirasawa", 15, 154.0},
	{"Nao Okuda", 17, 158.0},
}
