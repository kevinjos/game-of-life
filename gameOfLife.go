package main

import (
  "fmt"
  "math/rand"
  "time"
  "strings"
)

const X, Y int = 80, 40

type Light struct {
  status uint8 //0: Dead, 1: Born, 2: Dying, 3: Alive
}

type Board struct {
  matrix [Y][X]Light
}

func (b *Board) Print(i int) {
  for i := range b.matrix {
    for _, x := range b.matrix[i] {
      if x.status == 3 {
        fmt.Print("*")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n")
  }
  fmt.Println("Gen", i, strings.Repeat("~", X-6))
}

func (b *Board) Initialize() {
  seed := time.Now().UnixNano()
  source := rand.NewSource(seed)
  prng := rand.New(source)
  for i := 0; i < X ; i++ {
    for j := 0; j < Y; j++ {
      if prng.Int() % 8 == 0 {
        b.matrix[j][i].status = 3
      }
    }
  }
}

func (b *Board) Generation() {
  trans := func(i, j int) (uint8) {
    r := func(l Light) (uint8) {
      if l.status > 1 {
        return 1
      } else {
        return 0
      }
    }
    var nc uint8 = 0
    //Edge cases
    if i == X - 1 && j == Y - 1 {
      nc += r(b.matrix[j][i-1])
      nc += r(b.matrix[j-1][i-1])
      nc += r(b.matrix[j-1][i])
    } else if i == 0 && j == 0 {
      nc += r(b.matrix[j][i+1])
      nc += r(b.matrix[j+1][i+1])
      nc += r(b.matrix[j+1][i])
    } else if i == 0 && j == Y -1 {
      nc += r(b.matrix[j][i+1])
      nc += r(b.matrix[j-1][i+1])
      nc += r(b.matrix[j-1][i])
    } else if i == X - 1 && j == 0 {
      nc += r(b.matrix[j][i-1])
      nc += r(b.matrix[j+1][i-1])
      nc += r(b.matrix[j+1][i])
    } else if i == 0 {
      nc += r(b.matrix[j][i+1])
      nc += r(b.matrix[j+1][i+1])
      nc += r(b.matrix[j-1][i+1])
      nc += r(b.matrix[j+1][i])
      nc += r(b.matrix[j-1][i])
    } else if j == 0 {
      nc += r(b.matrix[j][i+1])
      nc += r(b.matrix[j+1][i+1])
      nc += r(b.matrix[j][i-1])
      nc += r(b.matrix[j+1][i-1])
      nc += r(b.matrix[j+1][i])
    } else if i == X - 1 {
      nc += r(b.matrix[j][i-1])
      nc += r(b.matrix[j+1][i-1])
      nc += r(b.matrix[j-1][i-1])
      nc += r(b.matrix[j+1][i])
      nc += r(b.matrix[j-1][i])
    } else if j == Y - 1 {
      nc += r(b.matrix[j][i+1])
      nc += r(b.matrix[j-1][i+1])
      nc += r(b.matrix[j][i-1])
      nc += r(b.matrix[j-1][i-1])
      nc += r(b.matrix[j-1][i])
    } else {
      //Base case
      nc += r(b.matrix[j][i+1])
      nc += r(b.matrix[j+1][i+1])
      nc += r(b.matrix[j-1][i+1])
      nc += r(b.matrix[j][i-1])
      nc += r(b.matrix[j+1][i-1])
      nc += r(b.matrix[j-1][i-1])
      nc += r(b.matrix[j+1][i])
      nc += r(b.matrix[j-1][i])
    }
    return nc
  }

  var nc uint8
  for j := range b.matrix {
    for i, x := range b.matrix[j] {
      nc = trans(i, j)
      if x.status == 0 && nc == 3 {
        b.matrix[j][i].status = 1
      } else if x.status == 3 && (nc < 2 || nc > 3) {
        b.matrix[j][i].status = 2
      } 
    }
  }

  for j := range b.matrix {
    for i, x := range b.matrix[j] {
      if x.status == 1 {
        b.matrix[j][i].status = 3
      } else if x.status == 2 {
        b.matrix[j][i].status = 0
      }
    }
  }
}

func main() {
  b := new(Board)
  b.Initialize()
  for i := 0; i > -1; i++ {
    b.Generation()
    b.Print(i)
    time.Sleep(500 * time.Millisecond)
  }
}
