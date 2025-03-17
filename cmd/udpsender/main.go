package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const addr = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("can't bind to the address: %v", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("can't setup the connection: %v", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("can't read the line: %v", err)
			return
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatalf("can't write to the UDP connection: %v", err)
			return
		}
	}

}
