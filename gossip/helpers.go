package gossip

import (
	"errors"
)

func SendRequest(endpoint string, args []byte) ([]byte, error) {
	// Initialize the request with a leading byte.
	request := []byte{0x00}

	// Append the endpoint and a separator byte to the request.
	request = append(request, endpoint...)
	request = append(request, 0x00)

	// Append the arguments to the request.
	request = append(request, args...)

	// Calculate the size of the request.
	size := len(request)

	// Check if the size exceeds the maximum value for a 32-bit unsigned integer (4,294,967,295).
	// If it does, return an error.
	if size > 4294967295 {
		return nil, errors.New("request too big")
	}

	// Prepare a 4-byte slice to hold the size of the request.
	size_header := make([]byte, 4)

	// Encode the size into the size_header using bit shifting.
	// size_header[0] holds the most significant byte,
	// and size_header[3] holds the least significant byte.
	size_header[0] = byte((size >> 24) & 0xFF)
	size_header[1] = byte((size >> 16) & 0xFF)
	size_header[2] = byte((size >> 8) & 0xFF)
	size_header[3] = byte(size & 0xFF)

	// Here you would typically continue with the logic for sending the request,
	// and then return the response and any error that occurs.

	// This is a placeholder return. Replace it with actual request sending logic.
	return append(size_header, request...), nil
}

func BareEncode(args ...[]byte) ([]byte, error) {
	if len(args)%3 != 0 {
		return nil, errors.New("INVALID: Amount of byte arrays provided insufficient")
	}

	request := []byte{}

	for i := 0; i < len(args)/3; i++ {

		varn := args[(1+(i*((1%(i+1))*3)))-1]
		vart := args[(2+(i*((1%(i+1))*3)))-1]
		varv := args[(3+(i*((1%(i+1))*3)))-1]

		request = append(request, varn...)
		request = append(request, 0x00)
		request = append(request, vart...)
		request = append(request, 0x00)
		request = append(request, varv...)

		if string(vart) == "str" || string(vart) == "xdr" {
			request = append(request, 0x00)
		}

	}

	request = append(request, 0x04)

	return request, nil
}
