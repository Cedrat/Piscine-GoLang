package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func giveAnswer(incoming string, client net.Conn) {

	var bjr []byte

	go func() {
		switch incoming {
		case "fr":
			bjr = []byte("Quarante-deux\n")
		case "en":
			bjr = []byte("Forty-Two\n")
		case "kr":
			bjr = []byte("사십이\n")
		default:
			bjr = []byte("Unknown\n")
		}
		for i := 0; i < 5; i++ {
			client.Write(bjr)
			time.Sleep(time.Second)
		}
	}()

}

func main() {
	nb_client := 0
	port := ":8000"
	if len(os.Args) == 3 || len(os.Args) == 2 {
		if os.Args[1] == "-p" {
			port = ":" + os.Args[2]
		} else {
			fmt.Println("Argument error, try -p PORT")
		}
	}
	listener, err := net.Listen("tcp", port)
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
			for {
				nb, err := conn.Read(recvData)
				if err != nil && err.Error() == "EOF" {
					nb_client--
					break
				}
				if nb > 0 {
					giveAnswer(string(recvData[:nb-1]), conn)
				}
			}
		}()
	}
}
