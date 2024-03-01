package tests

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

// Server starts a TCP server that echoes back any received message.
func Server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Received message: ", string(message))
	conn.Write([]byte("message\n"))
}

// Client connects to the TCP server and sends a message.
func Client(message string) (response string, err error) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}

	buffer, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}

	return buffer, nil
}

// TestServerClient tests the server and client interaction.
func TestServerClient(t *testing.T) {
	go Server() // Start the server in a goroutine

	// Give the server a moment to start
	time.Sleep(time.Second)

	// Test data
	message := "Hello, Server!\n"
	expectedResponse := "Hello, Server!\n"

	response, err := Client(message)
	if err != nil {
		t.Fatalf("Failed to receive response from server: %v", err)
	}

	if response != expectedResponse {
		t.Errorf("Expected response '%s', got '%s'", expectedResponse, response)
	}
}
