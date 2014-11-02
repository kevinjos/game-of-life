package main

import (
	"encoding/gob"
	"github.com/kevinjos/game-of-life/board"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":1112")
	if err != nil {
		log.Printf("Error dialing server: %s", err)
	}
	defer conn.Close()

	dec := gob.NewDecoder(conn)

	for {
		var b board.Board
		err := dec.Decode(&b)
		if err != nil {
			log.Printf("Error decoding: %s", err)
		}
		b.Print()
	}
}
