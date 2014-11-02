package board

import (
  "fmt"
  "strings"
  "time"
  "math/rand"
)

const X, Y int = 40, 40

type Light struct {
	Status uint8 //0: Dead, 1: Born, 2: Dying, 3: Alive
}

type Board struct {
	Matrix [Y][X]Light
}

func (b *Board) Print() {
  s := make([]string, 1)
	for i := range b.Matrix {
		for _, x := range b.Matrix[i] {
			if x.Status == 3 {
        s = append(s, "*")
				//fmt.Print("*")
			} else {
        s = append(s, " ")
				//fmt.Print(" ")
			}
		}
    s = append(s, "\n")
		//fmt.Print("\n")
	}
  s = append(s, strings.Repeat("~", 2*X))
	//fmt.Println("Gen", i, strings.Repeat("~", X-6))
  fmt.Println(s)
}

func (b *Board) Initialize() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	prng := rand.New(source)
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if prng.Int() % 6 == 0 {
				b.Matrix[j][i].Status = 3
			}
		}
	}
}

func (b *Board) Generation() {
	trans := func(i, j int) uint8 {
		r := func(l Light) uint8 {
			if l.Status > 1 {
				return 1
			} else {
				return 0
			}
		}
		var nc uint8 = 0
		//Edge cases
		if i == X-1 && j == Y-1 {
			nc += r(b.Matrix[j][i-1])
			nc += r(b.Matrix[j-1][i-1])
			nc += r(b.Matrix[j-1][i])
		} else if i == 0 && j == 0 {
			nc += r(b.Matrix[j][i+1])
			nc += r(b.Matrix[j+1][i+1])
			nc += r(b.Matrix[j+1][i])
		} else if i == 0 && j == Y-1 {
			nc += r(b.Matrix[j][i+1])
			nc += r(b.Matrix[j-1][i+1])
			nc += r(b.Matrix[j-1][i])
		} else if i == X-1 && j == 0 {
			nc += r(b.Matrix[j][i-1])
			nc += r(b.Matrix[j+1][i-1])
			nc += r(b.Matrix[j+1][i])
		} else if i == 0 {
			nc += r(b.Matrix[j][i+1])
			nc += r(b.Matrix[j+1][i+1])
			nc += r(b.Matrix[j-1][i+1])
			nc += r(b.Matrix[j+1][i])
			nc += r(b.Matrix[j-1][i])
		} else if j == 0 {
			nc += r(b.Matrix[j][i+1])
			nc += r(b.Matrix[j+1][i+1])
			nc += r(b.Matrix[j][i-1])
			nc += r(b.Matrix[j+1][i-1])
			nc += r(b.Matrix[j+1][i])
		} else if i == X-1 {
			nc += r(b.Matrix[j][i-1])
			nc += r(b.Matrix[j+1][i-1])
			nc += r(b.Matrix[j-1][i-1])
			nc += r(b.Matrix[j+1][i])
			nc += r(b.Matrix[j-1][i])
		} else if j == Y-1 {
			nc += r(b.Matrix[j][i+1])
			nc += r(b.Matrix[j-1][i+1])
			nc += r(b.Matrix[j][i-1])
			nc += r(b.Matrix[j-1][i-1])
			nc += r(b.Matrix[j-1][i])
		} else {
			//Base case
			nc += r(b.Matrix[j][i+1])
			nc += r(b.Matrix[j+1][i+1])
			nc += r(b.Matrix[j-1][i+1])
			nc += r(b.Matrix[j][i-1])
			nc += r(b.Matrix[j+1][i-1])
			nc += r(b.Matrix[j-1][i-1])
			nc += r(b.Matrix[j+1][i])
			nc += r(b.Matrix[j-1][i])
		}
		return nc
	}

	var nc uint8
	for j := range b.Matrix {
		for i, x := range b.Matrix[j] {
			nc = trans(i, j)
			if x.Status == 0 && nc == 3 {
				b.Matrix[j][i].Status = 1
			} else if x.Status == 3 && (nc < 2 || nc > 3) {
				b.Matrix[j][i].Status = 2
			}
		}
	}

	for j := range b.Matrix {
		for i, x := range b.Matrix[j] {
			if x.Status == 1 {
				b.Matrix[j][i].Status = 3
			} else if x.Status == 2 {
				b.Matrix[j][i].Status = 0
			}
		}
	}
}


