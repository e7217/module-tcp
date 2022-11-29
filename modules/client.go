package modules

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func InitClient() {

	time.Sleep(time.Second * 5)
	clientHandleConnection()
}

func clientHandleConnection() {
	fmt.Println("client attemps connection...")

	client, err := net.Dial("tcp", "127.0.0.1:8888")

	fmt.Println("client connected")

	defer client.Close()

	if err != nil {
		log.Fatal(err)
	}
	wg.Add(2)

	go clientRecvHandler(client)
	go clientSendHandler(client)
	wg.Wait()
}

func clientRecvHandler(conn net.Conn) {
	data := make([]byte, 512)

	for {
		n, err := conn.Read(data)
		if err != nil {
			log.Fatal(err)
			continue
		}

		fmt.Println(string(data[0:n]))
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

func clientSendHandler(conn net.Conn) {
	for {
		_, err := conn.Write([]byte("ping\n"))
		if err != nil {
			log.Fatal(err)
			continue
		}
		time.Sleep(time.Second * 1)
	}
	wg.Done()

}
