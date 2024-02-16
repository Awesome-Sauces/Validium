package gossip

import (
	"encoding/binary"
	"fmt"
)

type Connection struct {
	m_args map[string]FunctionArgument
	write  []byte // Set by the connection.Write() used at the end of server process
}

type connError struct {
	arg string
}

func (e *connError) Error() string {
	return fmt.Sprintf("!HYDROGEN! _ ***%s*** _ PLEASE REFRENCE SOURCECODE FOR ANY ISSUES", e.arg)
}

func NewConnection(args []FunctionArgument) *Connection {
	m_args := make(map[string]FunctionArgument)

	for _, val := range args {
		m_args[val.Name] = val
	}

	return &Connection{
		m_args: m_args,
	}
}

func (conn Connection) GetRawValue(id string) ([]byte, error) {

	val, contains := conn.m_args[id]

	if contains {
		return val.Bytes, nil
	}

	return []byte{}, &connError{arg: fmt.Sprintf("404 -> BYTE ARRAY NOT FOUND WITH ID (%s)", id)}

}

func (conn Connection) Dump(id string) {
	delete(conn.m_args, id)
}

func (conn Connection) GetString(id string) (string, error) {
	val, contains := conn.m_args[id]

	if contains {
		return string(val.Bytes), nil
	}

	return "", &connError{arg: fmt.Sprintf("404 -> STRING NOT FOUND WITH ID (%s)", id)}
}

// Fetches var by ID then converts to signed 64
func (conn Connection) GetInt(id string) (int64, error) {
	val, contains := conn.m_args[id]

	if contains {
		return int64(binary.LittleEndian.Uint64(val.Bytes)), nil
	}

	return int64(0), &connError{arg: fmt.Sprintf("404 -> STRING NOT FOUND WITH ID (%s)", id)}
}

// Will write back once process ends
func (conn *Connection) Write(val []byte) {
	conn.write = val
}

func (conn Connection) GetWrite() []byte {
	return conn.write
}
