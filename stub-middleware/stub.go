package stub

import (
	"fmt"
	"net"
)

type Response struct {
	Message string
}

func (s *Response) GetStubFromService() (*Response, error) {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	conn.Write([]byte(fmt.Sprintf("Get|ServiceFootball")))

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	addr := string(buf[:n])

	conn, err = net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	conn.Write([]byte("Get|CurrentRound|PL|2021"))

	buf = make([]byte, 1024)

	n, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}

	s.Message = string(buf[:n])
	return s, nil
}
