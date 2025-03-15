package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Println("can't open the file")
	}

	defer file.Close()

	b1 := make([]byte, 8)

	for {
		n1, err := file.Read(b1)
		if err != nil && err != io.EOF {
			log.Println("can't read the file")
		}
		if n1 == 0 {
			break
		}

		fmt.Printf("read: %s\n", string(b1[:n1]))
	}
}
