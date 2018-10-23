package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	var listener, err = net.Listen("tcp", ":43594")

	if err != nil {
		log.Fatal("Failed to start server", err)
		return
	}

	log.Print("Listening for new connections")

	for {

		var conn, err = listener.Accept()

		if err != nil {
			log.Print("Failed to accept new connection", err)
			continue
		}

		log.Print("Accepted new connection", conn.LocalAddr())

		var client = &Client{
			connection: conn,
			reader: bufio.NewReader(conn),
			writer: bufio.NewWriter(conn),
			buffer: NewBuffer(256),
		}

		go client.Listen()
	}
}
