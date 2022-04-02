package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type Client struct {
	name   string
	socket net.Conn
}

func NewClient(name string, socket net.Conn) (*Client, error) {
	client := Client{name, socket}

	return &client, nil
}

func (client *Client) Name() string {
	return client.name
}

func SendMessageToAll(clients []Client, msg string) {
	for _, client := range clients {
		client.socket.Write([]byte(msg))
	}
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

	var clients []Client

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error encountered")
		}
		new_client, _ := NewClient(strconv.Itoa(nb_client), conn)
		nb_client++
		SendMessageToAll(clients, new_client.Name()+" participated.\n")
		clients = append(clients, *new_client)
		new_client.socket.Write([]byte("Connection Success, If you want to end, input EOF\n"))
		new_client.socket.Write([]byte("login : " + new_client.Name() + "\n"))
		fmt.Println("Nb client connected :", nb_client)
		go func() {
			recvData := make([]byte, 1024)
			defer conn.Close()
			for {
				nb, err := conn.Read(recvData)
				if err != nil && err.Error() == "EOF" {
					SendMessageToAll(clients, new_client.Name()+" left the room.")
					break
				}
				if nb > 0 {
					login := []byte("login " + new_client.Name() + ": ")
					SendMessageToAll(clients, string(login)+string(recvData[:nb]))
				}
			}
		}()
	}
}
