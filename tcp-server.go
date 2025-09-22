package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	fmt.Println("Server running on port 8080")

	for {
		cnn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(cnn)
	}
}

func handle(cnn net.Conn) {
	scanner := bufio.NewScanner(cnn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(cnn, "got you body %s", ln)
	}
	defer cnn.Close()
}
