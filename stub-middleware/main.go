package stub

import (
	"fmt"
	"net"
)

type Stub struct {
	Message string
}

func (s *Stub) GetStubFromService() error {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		return err
	}

	defer conn.Close()

	conn.Write([]byte(fmt.Sprintf("Get|ServiceFootball")))

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return err
	}

	addr := string(buf[:n])

	conn, err = net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	defer conn.Close()

	conn.Write([]byte("Get|CurrentRound|PL|2021"))

	buf = make([]byte, 1024)

	n, err = conn.Read(buf)
	if err != nil {
		return err
	}

	s.Message = string(buf[:n])
	return nil
}
