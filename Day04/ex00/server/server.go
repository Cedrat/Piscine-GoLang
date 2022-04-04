package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	nb_client := 0
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("error when launching server")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error encountered")
		}
		nb_client++
		fmt.Println("Nb client connected :", nb_client)
		go func() {
			recvData := make([]byte, 1024)
			defer conn.Close()
			hi := []byte("42\n")
			for {
				conn.SetReadDeadline(time.Now().Add(time.Second))

				conn.Write(hi)
				_, err = conn.Read(recvData)
				if err != nil && err.Error() == "EOF" {
					nb_client--
					break
				}
			}
		}()
	}
}
