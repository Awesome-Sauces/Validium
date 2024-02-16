package gossip

import (
	"log"
	"net"
)

func StartTCPServer() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	log.Println("Server listening on :8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting:", err.Error())
			continue
		}

		endpoint, args := DigestRequest((conn))
		variables := MapVariables(args)

		log.Println(endpoint)

		for key, val := range variables {
			log.Println(key, ":", val.Name, ":", val.Type, ":", val.Bytes)
		}
	}
}
