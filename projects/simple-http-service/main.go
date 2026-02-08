package main

import (
	"http-service/handlers"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s", addr)

	for {
		conn, err := ln.Accept()
		log.Println("connection", conn)
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handlers.HandleHttp9(conn)
	}
}
