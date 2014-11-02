package main

import (
	"encoding/gob"
	"github.com/kevinjos/game-of-life/board"
	"log"
	"net"
	"time"
)

func main() {
	b := new(board.Board)
	b.Initialize()
	l, err := net.Listen("tcp", ":1112")
	if err != nil {
		log.Printf("Error creating listener: %s", err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
		}
		defer conn.Close()
		go func(c net.Conn) {
			enc := gob.NewEncoder(c) // Will write to network.
			for {
				time.Sleep(300 * time.Millisecond)
				b.Generation()
				err := enc.Encode(b)
				if err != nil {
					log.Printf("Error encoding: %s", err)
					return
				}
			}
		}(conn)
	}
}
