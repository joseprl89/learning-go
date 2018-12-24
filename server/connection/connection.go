package connection

import "net"

func ReadFrom(conn net.Conn) (string, error) {
	bufferSize := 32
	buffer := make([]byte, bufferSize)
	var result []byte
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			return "", err
		} else {
			result = append(result, buffer[0:n]...)
		}

		if n < bufferSize {
			return string(result[:len(result)]), nil
		}
	}
}
