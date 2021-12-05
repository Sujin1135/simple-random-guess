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
		if result == int8(-1) {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		} else if result == int8(1) {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else {
			fmt.Println("정답입니다.")
			break
		}
	}
	fmt.Println("게임을 종료합니다.")
}
