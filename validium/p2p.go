package validium

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/Awesome-Sauces/Validium/gossip"
)

type HandleRequest struct {
}

func (handler HandleRequest) Call(conn *gossip.Connection) {
	fmt.Println("PING!")

	val, err := conn.GetString("address")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(val)

	conn.Write([]byte{0x60})
}

func Test() {
	start_server := func() {
		server := gossip.NewServer("localhost:8080")

		server.NewEndpoint("test", HandleRequest{})

		server.Listen()
	}

	go start_server()

	time.Sleep(1 * time.Second)

	server := gossip.NewServer("localhost:9090")

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(1234567891234567890))
	values, err := gossip.BareEncode([]byte("token"), []byte("i64"), b, []byte("name"), []byte("str"), []byte("xrp"), []byte("address"), []byte("str"), []byte("0xA1FC67E"), []byte("send_amount"), []byte("i64"), b)

	if err != nil {
		log.Fatal(err)
	}

	dialval, err := gossip.SendRequest("test", values)

	if err != nil {
		log.Fatal(err)
	}

	retval, err := server.Dial("localhost:8080", dialval)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(retval[0])
}
