package main

import "fmt"

func phepToan(number1, number2 int) (int, int, int, float32) {
	if number2 == 0 {
		number2 = 5
	}
	sum := number1 + number2
	hieu := number1 - number2
	tich := number1 * number2
	thuong := float32(number1) / float32(number2)
	return sum, hieu, tich, thuong
}

func countDown(number int) {
	fmt.Println(number)
	if number > 0 {
		countDown(number - 1)
	}
}

func main() {
	//cong, tru, nhan, chia := phepToan(1, 10)
	//fmt.Println("Cong: ", cong)
	//fmt.Println("Tru: ", tru)
	//fmt.Println("Nhan: ", nhan)
	//fmt.Println("Chia: ", chia)

	// De quy
	countDown(10)
}
