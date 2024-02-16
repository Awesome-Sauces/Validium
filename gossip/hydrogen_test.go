package gossip

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

type Node struct {
	ID    int
	Peers map[int]chan []byte
}

func (n Node) Gossip(transaction []byte) {

}

func TestOpenTrustProtocol(t *testing.T) {

}

func TestSizeCalculation(t *testing.T) {
	buffer := []byte{
		0x00,
		0x00,
		0x00,
		0x15,
	}

	size := binary.BigEndian.Uint32(buffer[:4])
	log.Printf("Interpreted as 0x00000015, the size is: %d", size)

	// To demonstrate the original scenario:
	buffer265 := []byte{
		0x00,
		0x00,
		0x01,
		0x05,
	}

	size265 := binary.BigEndian.Uint32(buffer265[:4])
	log.Printf("Interpreted as 0x00000105, the size is: %d", size265)

	fmt.Println("Length of buffer:", len(buffer))
	fmt.Println("Length of buffer265:", len(buffer265))

}

func TestMapVariables(t *testing.T) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(1234567891234567890))
	values, err := BareEncode([]byte("token"), []byte("i64"), b, []byte("name"), []byte("str"), []byte("xrp"), []byte("address"), []byte("str"), []byte("0xA1FC67E"), []byte("send_amount"), []byte("i64"), b)

	if err != nil {
		log.Fatal(err)
	}
	pass_back := MapVariables(values)

	for key, val := range pass_back {
		log.Println(key, ":", val.Name, ":", val.Type, ":", val.Bytes)
	}
}

func TestServerCore(t *testing.T) {
	go StartTCPServer()

	time.Sleep(1 * time.Second)

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(1234567891234567890))
	values, err := BareEncode([]byte("token"), []byte("i64"), b, []byte("name"), []byte("str"), []byte("xrp"), []byte("address"), []byte("str"), []byte("0xA1FC67E"), []byte("send_amount"), []byte("i64"), b)

	if err != nil {
		log.Fatal(err)
	}

	request, err := SendRequest("StartPropagation", values)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	conn.Write(request)
	conn.Close()

	select {}
}
