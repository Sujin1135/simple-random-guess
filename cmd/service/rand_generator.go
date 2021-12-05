package service

import (
	"fmt"
	"math/rand"
	"time"
)

type RandGenerator struct {
	Num int8
}

func NewRandGenerator() *RandGenerator {
	rand.Seed(time.Now().UnixNano())
	return &RandGenerator{Num: int8(rand.Intn(100) + 1)}
}

func (g *RandGenerator) GuessNum(num int8) int8 {
	if g.Num == num {
		fmt.Println("정답입니다.")
		return int8(0)
	} else if g.Num > num {
		fmt.Println("입력하신 숫자가 더 작습니다.")
		return -1
	} else {
		fmt.Println("입력하신 숫자가 더 큽니다.")
		return 1
	}
}
