package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	log.Println("Server is listening...")

	clients := make(map[string]string)

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			return
		}

		clientAddr := addr.String()
		message := string(buffer[:n])

		if strings.Contains(message, "Get") {
			serv := strings.Split(message, "|")[1]
			_, err := conn.WriteToUDP([]byte(clients[serv]), addr)
			if err != nil {
				log.Println("Error sending message to stub")
			}
		}

		if strings.Contains(message, "Add") {
			log.Println("Client connected:", clientAddr)
			add := strings.Split(message, "|")
			clients[add[1]] = add[2]
		}

		if strings.Contains(message, "Delete") {
			log.Println("Client disconnected:", clientAddr)
			rem := strings.Split(message, "|")
			delete(clients, rem[1])
		}

		log.Printf("Received from client: %s message: %s\n", clientAddr, message)
		log.Println("Client list:", clients)
	}
}
