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

	ch := getLinesChannel(file)

	for line := range ch {
		fmt.Printf("read: %s\n", line)
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
			// read the file
			bytesRead, err := f.Read(buffer)
			if err != nil && err != io.EOF {
				log.Printf("can't read the file: %v", err)
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
