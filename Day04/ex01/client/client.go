package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":8000"
	if len(os.Args) == 3 || len(os.Args) == 2 {
		if os.Args[1] == "-p" {
			port = ":" + os.Args[2]
		} else {
			fmt.Println("Argument error, try -p PORT")
		}
	}
	conn, err := net.Dial("tcp", port)
	if err != nil {
		fmt.Println("Error when initializing client ")
		return
	}
	var reply = make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, text)
		go func() {
			for {
				nb, _ := conn.Read(reply)
				if nb > 0 {
					fmt.Println(string(reply[:nb-1]))
				}
			}
		}()

	}
}
