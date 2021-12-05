package service

import (
	"math/rand"
	"time"
)

type RandGenerator struct {
	Num int8
}

func NewRandGenerator () *RandGenerator {
	rand.Seed(time.Now().UnixNano())
	return &RandGenerator{Num: int8(rand.Intn(100) + 1)}
}

func (g *RandGenerator) GuessNum(num int8) int8 {
	if g.Num == num {
		return int8(0)
	} else if g.Num > num {
		return -1
	} else {
		return 1
	}
}
