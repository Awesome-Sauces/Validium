package gossip

import (
	"encoding/binary"
	"log"
	"net"
)

const (
	NULL        = 0
	String      = 1
	Signed32    = 2
	Unsigned32  = 3
	Signed64    = 4
	Unsigned64  = 5
	Signed128   = 6
	Unsigned128 = 7
)

type FunctionArgument struct {
	Name  string
	Type  string
	Bytes []byte
}

func min(x int, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x int, y int) int {
	if x < y {
		return y
	}

	return x
}

func DigestRequest(conn net.Conn) (string, []byte) {
	defer conn.Close()

	buffer := make([]byte, 4) // Buffer to read 4 bytes
	_, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return "nil", nil
	}

	// Convert the first 4 bytes to a uint32 using big-endian format
	size := binary.BigEndian.Uint32(buffer[:4])

	buffer = make([]byte, size)
	_, err = conn.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err.Error())
		return "FAILED TO READ BYTES (MIGHT HAVE BYTE OVERFLOW/UNDERFLOW)", nil
	}

	b_endpoint := []byte{}

	t_buffer := []byte{}

	t_activation := false

	for i := range buffer {
		val := buffer[min(max(len(buffer)-1, 0), i+1)]
		if !t_activation {
			if val == 0 && !t_activation {
				t_activation = true
			} else if !t_activation {
				b_endpoint = append(b_endpoint, val)
			}
		} else {
			t_buffer = append(t_buffer, val)
		}

	}

	return string(b_endpoint), t_buffer
}

func MapVariables(data []byte) []FunctionArgument {
	var retval []FunctionArgument

	i := 0
	for i < len(data) && data[i] != 0x04 { // Stop if the termination byte (0x04) is encountered
		varName, varType, varValue := []byte{}, []byte{}, []byte{}

		// Parse variable name
		for i < len(data) && data[i] != 0x00 {
			varName = append(varName, data[i])
			i++
		}
		i++ // Skip null byte

		// Parse variable type
		for i < len(data) && data[i] != 0x00 {
			varType = append(varType, data[i])
			i++
		}
		i++ // Skip null byte

		// Parse variable value based on type
		switch string(varType) {
		case "i32", "u32":
			if i+4 <= len(data) {
				varValue = append(varValue, data[i:i+4]...)
				i += 4
			}
		case "i64", "u64":
			if i+8 <= len(data) {
				varValue = append(varValue, data[i:i+8]...)
				i += 8
			}
		case "i128", "u128":
			if i+16 <= len(data) {
				varValue = append(varValue, data[i:i+16]...)
				i += 16
			}
		case "str", "xdr":
			for i < len(data) && data[i] != 0x00 {
				varValue = append(varValue, data[i])
				i++
			}
			i++ // Skip null byte for string
		}

		// Append the parsed variable to the result
		if len(varName) > 0 && len(varType) > 0 {
			retval = append(retval, FunctionArgument{Name: string(varName), Type: string(varType), Bytes: varValue})
		}
	}

	return retval
}
