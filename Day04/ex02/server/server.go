package main

import (
	"fmt"
	"log"
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

func SetLogger() log.Logger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return *logger
}

func ClientRoutines(clients []Client, curr_client Client, logger log.Logger) {
	go func() {
		recvData := make([]byte, 1024)
		defer curr_client.socket.Close()
		for {
			nb, err := curr_client.socket.Read(recvData)
			if err != nil && err.Error() == "EOF" {
				SendMessageToAll(clients, curr_client.Name()+" left the room.\n")
				logger.Print("logout : ", curr_client.socket.LocalAddr().String(), " ( login ", curr_client.Name(), " )")
				break
			}
			if nb > 0 {
				login := []byte("login " + curr_client.Name() + ": ")
				SendMessageToAll(clients, string(login)+string(recvData[:nb]))
			}
		}
	}()
}

func MainLoop(listener net.Listener, logger log.Logger) {
	var clients []Client
	nb_client := 0
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
		logger.Print("login : ", conn.LocalAddr().String(), " ( login ", new_client.Name(), " )")
		ClientRoutines(clients, *new_client, logger)
	}
}
func main() {
	logger := SetLogger()

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

	logger.Print("Server starting succeed ! ")
	MainLoop(listener, logger)

}
