package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(number int) int {
	if number == 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}

func fibo(number int) int {
	term1, term2 := 0, 1
	if number == 0 {
		fmt.Println("Day fibonacci:", term1)
		fmt.Println("Tong day fibonacci:", term1)
	} else if number == 1 {
		fmt.Println("Day fibonacci:", term2)
		fmt.Println("Tong day fibonacci:", term1+term2)
	} else {
		total := term1 + term2
		fmt.Print("Day fibonacci: ", term1, term2)

		for i := 3; i <= number; i++ {
			term3 := term1 + term2

			fmt.Printf(" %d", term3)

			total += term3

			term1, term2 = term2, term3
		}

		fmt.Println()
		fmt.Println("Tong day fibonacci: ", total)
	}

	return 0
}

func readInt(prompt string) (int, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	input = strings.TrimSpace(input)
	numb, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("error parsing input: %v", err)
	}

	return numb, nil
}

func main() {
	for {
		fmt.Println("-=-=-=-=-= Menu Chuc Nang -=-=-=-=-=")
		fmt.Println("[1] Tinh tong day so N phan tu")
		fmt.Println("[2] Hien thi va tinh tong day so Fibonacci")
		fmt.Println("[0] Thoat chuong trinh")
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

		var choice int
		for {
			var err error
			choice, err = readInt("Vui long nhap so chuc nang muon thuc hien: ")
			if err != nil || choice < 0 {
				fmt.Println("⛔️ Vui long nhap mot so nguyen")
			} else {
				break
			}
		}

		switch choice {
		case 1:
			var numb int
			for {
				var err error
				numb, err = readInt("Nhap N > 0: ")
				if err != nil || numb <= 0 {
					fmt.Println("⛔ Vui long nhap vao la so nguyen duong N")
				} else {
					break
				}
			}
			result := sum(numb)
			fmt.Printf("Tong tu 1 den %d co ket qua la: %d \n", numb, result)
		case 2:
			var numb int
			for {
				var err error
				numb, err = readInt("Nhap so luong day finonacci: ")
				if err != nil || numb < 0 {
					fmt.Println("⛔ So luong day finonacci phai lon hon hoac bang 0")

				} else {
					break
				}
			}
			fibo(numb)
		case 0:
			fmt.Println("✅ Cam on ban da su dung chuong trinh. Bye!")
			return
		default:
			fmt.Println("🛑 Vui long chon dung so chuc nang muon thuc hien")
		}
	}
}
