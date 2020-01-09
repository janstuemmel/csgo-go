package main

import (
	"bufio"
	"net"
	"os"
)

func main() {

	var file *os.File
	var err error

	address := "0.0.0.0:1234"

	if len(os.Args) < 2 {
		file = os.Stdin
	} else {
		file, err = os.Open(os.Args[1])
	}

	r := bufio.NewReader(file)

	l, _, err := r.ReadLine()

	conn, err := net.Dial("udp", address)

	defer conn.Close()

	for err == nil {
		conn.Write(l)
		l, _, err = r.ReadLine()
	}
}
