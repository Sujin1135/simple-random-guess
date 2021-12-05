package main

import (
	"fmt"
	"guess_rand_num/cmd/service"
)

func main() {
	randGenerator := service.NewRandGenerator()
	var guessNum int8

	for {
		fmt.Println("예상되는 숫자를 입력하세요")
		fmt.Scan(&guessNum)
		result := randGenerator.GuessNum(guessNum)
		if result == int8(0) {
			break
		}
	}
	fmt.Println("게임을 종료합니다.")
}
