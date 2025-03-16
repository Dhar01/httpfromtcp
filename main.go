package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("can't listen: %v\n", err)
	}

	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("incoming connection error: %v\n", err)
		}

		fmt.Println("connection has been accepted from:", conn.RemoteAddr())

		ch := getLinesChannel(conn)

		for line := range ch {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	// create a channel
	ch := make(chan string)

	go func() {
		// creating buffer to read 8 bytes at a time
		buffer := make([]byte, 8)
		currentLine := ""

		// closing channel when goroutine ends
		defer close(ch)
		defer f.Close()
		for {
			// read the connection (previously file)
			bytesRead, err := f.Read(buffer)
			if err != nil && err != io.EOF {
				log.Printf("can't read the connection: %v\n", err)
				return
			}

			if err == io.EOF {
				break
			}

			// add the new bytes to the current line
			currentLine += string(buffer[:bytesRead])

			if strings.Contains(currentLine, "\n") {
				parts := strings.SplitN(currentLine, "\n", 2)

				ch <- parts[0]

				if len(parts) > 1 {
					currentLine = parts[1]
				} else {
					currentLine = ""
				}
			}
		}

		if currentLine != "" {
			ch <- currentLine
		}
	}()

	return ch
}
