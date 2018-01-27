package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
)

func server(ln net.Listener, quotes *[]string) error {
	for {
		c, err := ln.Accept()
		if err != nil {
			return err
		}

		go func() {
			if err := handler(c, quotes); err != nil {
				log.Printf("Error while handling request from %v: %v", c.RemoteAddr(), err)
			}
		}()
	}
}

func handler(c net.Conn, quotes *[]string) error {
	defer c.Close()

	a := c.RemoteAddr()
	log.Println("New connection: " + a.String())

	r := rand.Intn(len(*quotes) - 1)
	c.Write([]byte((*quotes)[r]))
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("usage: port, file")
		return
	}

	bs, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatal("Error while reading file: %v", err)
	}

	q := strings.Split(string(bs), "\n")

	ln, err := net.Listen("tcp", ":"+args[0])
	if err != nil {
		log.Fatal("Error while creating listener: %v", err)
	}

	log.Println("Starting Server")
	server(ln, &q)
}
