package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
)

func randomQuote() ([]byte, error) {
	bs, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := strings.Split(string(bs), "\n")
	q = q[:len(q)-1]
	r := rand.Intn(len(q))

	return []byte(q[r]), nil
}

func server(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handler(c)
	}
}

func handler(c net.Conn) {
	defer c.Close()

	a := c.RemoteAddr()
	fmt.Println("New connection: " + a.String())

	q, err := randomQuote()
	if err != nil {
		log.Println(err)
		return
	}
	c.Write(q)
}

func main() {
	a := os.Args[1:]
	if len(a) < 2 {
		fmt.Println("Incorrect arguments. The correct order is");
		fmt.Println("usage: port, file")
		return
	}

	fmt.Println("Starting Server")

	server(string(a[0]))
}
