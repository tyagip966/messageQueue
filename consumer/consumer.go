package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"
)

var (
	connected     bool
	connectedSync sync.Mutex
)

func main() {
	log.Println("Consumer started...")
	for {
		connectedSync.Lock()
		alreadyConnected := connected
		connectedSync.Unlock()
		if !alreadyConnected {
			conn, err := net.Dial("tcp", "127.0.0.1:8000")
			if err != nil {
				log.Println(err.Error())
				time.Sleep(time.Duration(5) * time.Second)
				continue
			}
			log.Println(conn.RemoteAddr().String() + ": connected")
			connectedSync.Lock()
			connected = true
			connectedSync.Unlock()
			go receiveData(conn)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

//receiveData ... receiving data from the queuing service and printing.
func receiveData(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(conn.RemoteAddr().String() + ": disconnected")
			conn.Close()
			connectedSync.Lock()
			connected = false
			connectedSync.Unlock()
			log.Println(conn.RemoteAddr().String() + ": end receiving data")
			return
		}
		log.Print(message)
	}
}