package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var (
	connected     bool
	connectedSync sync.Mutex
)

func main() {
	log.Println("Client started...")
	for {
		connectedSync.Lock()
		alreadyConnected := connected
		connectedSync.Unlock()
		if !alreadyConnected {
			conn, err := net.Dial("tcp", "127.0.0.1:8001")
			if err != nil {
				log.Println(err.Error())
				time.Sleep(time.Duration(5) * time.Second)
				continue
			}
			log.Println(conn.RemoteAddr().String() + ": connected")
			connectedSync.Lock()
			connected = true
			connectedSync.Unlock()
			go sendData(conn)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

//sendData ... sending data to the queuingService.
func sendData(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		_, err := fmt.Fprintf(conn, message+"\n")

		if err != nil {
			log.Println(conn.RemoteAddr().String() + ": end sending data")
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
