package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Println("can't open the file")
	}

	defer file.Close()

	// for saving split line
	var parts []string

	buffer := make([]byte, 8)
	currentLine := ""

	for {
		// read the file
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Println("can't read the file")
		}

		if bytesRead == 0 {
			break
		}

		// add the new bytes to the current line
		currentLine += string(buffer[:bytesRead])

		if strings.Contains(currentLine, "\n") {
			parts = strings.SplitN(currentLine, "\n", 2)
			fmt.Printf("read: %s\n", parts[0])
			if len(parts) > 1 {
				currentLine = parts[1]
			} else {
				currentLine = ""
			}
		}
	}

	fmt.Printf("read: %s\n", currentLine)
}
