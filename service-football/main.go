package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"service-football/internal/routes"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("Error when init tcp connection", err)
		return
	}

	addr := ln.Addr().(*net.TCPAddr)
	log.Printf("Server is running on %s:%d\n", addr.IP, addr.Port)

	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Println("Error when connect to server", err)
		return
	}

	defer conn.Close()

	log.Println("Sending message to name server to register football service")
	conn.Write([]byte(fmt.Sprintf("Add|ServiceFootball|%s:%d", addr.IP, addr.Port)))

	c, err := routes.NewClient()
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error when accept connection", err)
			return
		}

		log.Println("Getting new connection:", conn.RemoteAddr())

		go handleConnection(conn, *c)
	}

}

func handleConnection(conn net.Conn, c routes.Client) {
	defer func() {
		log.Println("Connection closed:", conn.RemoteAddr())
		conn.Close()
	}()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Cannot read the data:", err)
			return
		}

		req := strings.Split(string(buf[:n]), "|")
		log.Print("Receiving message from", conn.RemoteAddr(), req[0], req[1], req[2], req[3])

		round, err := c.GetCurrentRound(req[2], req[3])
		if err != nil {
			panic(err)
		}

		matchDay, err := c.GetEvents(req[2], req[3], round.Response[0])
		if err != nil {
			panic(err)
		}

		jsonData, err := json.Marshal(matchDay)
		if err != nil {
			log.Println("Error marshalling data:", err)
			return
		}

		log.Println("Responding to client", conn.RemoteAddr(), "with data", round.Response[0])

		conn.Write(jsonData)
	}
}
