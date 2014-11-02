package main

import (
  "net"
  "log"
  "encoding/gob"
	"time"
  "github.com/kevinjos/game-of-life/board"
)

func main() {
	b := new(board.Board)
	b.Initialize()
  l, err := net.Listen("tcp", ":1112")
  if err != nil {
    log.Printf("Error creating listener: %s", err)
  }
	for {
    conn, err := l.Accept()
    if err != nil {
      log.Printf("Error accepting connection: %s", err)
    }
    defer l.Close()
    enc := gob.NewEncoder(conn) // Will write to network.
    go func (gob.Encoder) {
      for {
        time.Sleep(300 * time.Millisecond)
        b.Generation()
        b.Print()
        err := enc.Encode(b)
        if err != nil {
          log.Printf("Error encoding: %s", err)
        }
      }
    }(*enc)
	}
}
