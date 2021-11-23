package consumerClient

import (
	"fmt"
	"io"
	"net"
	"time"
)

type Consumers struct {
	Connection []net.Conn
}

func NewConsumers() *Consumers {
	return &Consumers{}
}

//AddNewConnection ... add new connection to the consumer list.
func (c *Consumers) AddNewConnection(newCon net.Conn) {
	c.Connection = append(c.Connection, newCon)
}

//RemoveClosedConnection ... removed closed connection from the Consumer array.
func (c *Consumers) RemoveClosedConnection(closedConn net.Conn) {
	for i := 0; i < len(c.Connection); i++ {
		n := c.Connection[i]
		if n == closedConn {
			c.Connection[i], c.Connection[len(c.Connection)-1] = c.Connection[len(c.Connection)-1], c.Connection[i]
			c.Connection = c.Connection[:len(c.Connection)-1]
			i--
		}
	}
}

//ConnectToConsumers ... function to open port to make connection with consumers
func ConnectToConsumers(consumer *Consumers) {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting socket server: " + err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error listening to client: " + err.Error())
			continue
		}
		consumer.AddNewConnection(conn)
		fmt.Println(conn.RemoteAddr().String() + ": consumer connected")
		go handleConnection(conn,consumer)
	}
}

//handleConnection ... this function is for handling the active consumers.
func handleConnection(conn net.Conn,consumer *Consumers) {
	defer conn.Close()
	notify := make(chan error)

	go func() {
		buf := make([]byte, 1024)
		for {
			_, err := conn.Read(buf)
			if err != nil {
				notify <- err
				consumer.RemoveClosedConnection(conn)
				return
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			if io.EOF == err {
				return
			}
		case <-time.After(time.Second * 1):
		}
	}
}