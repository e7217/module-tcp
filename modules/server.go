package modules

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func InitServer() {
	fmt.Println("server running 8888 port")
	socket, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	defer socket.Close()

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	wg1.Add(2)

	go recvHandler(conn)
	go sendHandler(conn)
	wg1.Wait()
}

func recvHandler(conn net.Conn) {
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[0:n]))
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

func sendHandler(conn net.Conn) {
	for {
		conn.Write([]byte("pong\n"))
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
