package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HinhChuNhat struct {
	d float32 `desc:"Chieu dai hinh chu nhat"`
	r float32 `desc:"Chieu rong hinh chu nhat"`
}

// Chu vi hinh chu nhat
// - Cong thuc: (d + r) * 2
// @return float32
func (hcn *HinhChuNhat) chuVi() float32 {
	return (hcn.d + hcn.r) * 2
}

// Dien tich hinh chu nhat
// - Cong thuc: d * r
// @return float32
func (hcn *HinhChuNhat) dienTich() float32 {
	return hcn.d * hcn.r
}

func readFloat(prompt string) float32 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Printf("❌ Lỗi: vui lòng nhập số hợp lệ: %v \n", err)
			continue
		}

		value, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Printf("❌ Lỗi: chuyển đổi kiểu dữ liệu: %v \n", err)
			continue
		}

		if value <= 0 {
			fmt.Println("⚠️ Giá trị phải lớn hơn 0.")
			continue
		}

		return float32(value)

	}
}

func main() {
	// Bai 1
	d := readFloat("Nhập chiều dài d > 0: ")
	r := readFloat("Nhập chiều rộng r > 0: ")
	hinh1 := HinhChuNhat{
		d: d,
		r: r,
	}

	cv := hinh1.chuVi()
	dt := hinh1.dienTich()
	fmt.Printf("Chu vi = %.2f \n", cv)
	fmt.Printf("Dien tich = %.2f \n", dt)

}
