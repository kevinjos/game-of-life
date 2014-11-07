package board

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const X, Y int = 80, 40

type Light struct {
	Status uint8 //0: Dead, 1: Born, 2: Dying, 3: Alive
}

type Board struct {
	Matrix [Y][X]Light
}

func (b *Board) Print() {
	s := ""
	for i := range b.Matrix {
		for _, x := range b.Matrix[i] {
			if x.Status == 3 {
				s += "*"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	s += strings.Repeat("~", X)
	fmt.Println(s)
}

func (b *Board) Initialize() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	prng := rand.New(source)
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if prng.Int()%6 == 0 {
				b.Matrix[j][i].Status = 3
			}
		}
	}
}

func (b *Board) Generation(topology string) {
	trans := func(i, j int, topology string) uint8 {
		r := func(l Light) uint8 {
			if l.Status > 1 {
				return 1
			} else {
				return 0
			}
		}
		var nc uint8 = 0
		if strings.EqualFold(topology, "closed") {
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
		} else if strings.EqualFold(topology, "wrapped") {
			nc += r(b.Matrix[(j+Y)%Y][(i+1+X)%X])
			nc += r(b.Matrix[(j+1+Y)%Y][(i+1+X)%X])
			nc += r(b.Matrix[(j-1+Y)%Y][(i+1+X)%X])
			nc += r(b.Matrix[(j+Y)%Y][(i-1+X)%X])
			nc += r(b.Matrix[(j+1+Y)%Y][(i-1+X)%X])
			nc += r(b.Matrix[(j-1+Y)%Y][(i-1+X)%X])
			nc += r(b.Matrix[(j+1+Y)%Y][(i+X)%X])
			nc += r(b.Matrix[(j-1+Y)%Y][(i+X)%X])
		}
		return nc
	}

	var nc uint8
	for j := range b.Matrix {
		for i, x := range b.Matrix[j] {
			nc = trans(i, j, topology)
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
