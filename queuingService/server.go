package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"queuingService/consumerClient"
	"queuingService/jsonQueue"
	"strings"
	"time"
)

var consumer *consumerClient.Consumers
var job *jsonQueue.Docs

func main() {
	consumer = consumerClient.NewConsumers()
	job = jsonQueue.NewJob()
	go connectToProducers()
	consumerClient.ConnectToConsumers(consumer)
}

//connectToProducers ... function to open port to make connection with producers
func connectToProducers() {
	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println("Error starting socket server: " + err.Error())
	}
	jobChan := make(chan *jsonQueue.Docs, 2)
	go newWorker(jobChan)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error listening to client: " + err.Error())
			continue
		}

		fmt.Println(conn.RemoteAddr().String() + ": producer connected")
		go receiveData(conn, jobChan)
	}
}

//receiveData ... inside this receiving data from producers and pushing to the job queue.
func receiveData(conn net.Conn, jobChan chan *jsonQueue.Docs) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": client disconnected")
			conn.Close()
			fmt.Println(conn.RemoteAddr().String() + ": end receiving data")
			return
		}
		var doc interface{}
		message = strings.Replace(message, "}", ",", len(message))
		message = message + "\"time_stamp\":" + "\"" + time.Now().String() + "\"}"
		_ = json.Unmarshal([]byte(message), &doc)
		job.AppendNewJob(doc)
		jobChan <- job
	}
}

//newWorker ... this worker takes the data from the job queue and send to the active consumers.
func newWorker(jobChan chan *jsonQueue.Docs) {
		for j := range jobChan {
			d := j.Pop()
			if d != nil {
				go sendDataToConsumer(d)
			}
		}
	}


//sendDataToConsumer ... sending data to the consumer.
func sendDataToConsumer(docs interface{}) {
	for _, conn := range consumer.Connection {
		bytesData, _ := json.Marshal(docs)
		_, err := fmt.Fprintf(conn, string(bytesData)+"\n")
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": end sending data")
			return
		}
	}
}
