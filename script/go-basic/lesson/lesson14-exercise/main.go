package main

import (
	"fmt"

	"trungem.com/hoc-golang/student"
	"trungem.com/hoc-golang/teacher"
	"trungem.com/hoc-golang/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("📚 CHUONG TRINH QUAN LY")
		fmt.Println("1. Quan ly sinh vien")
		fmt.Println("2. Quan ly giang vien")
		fmt.Println("3. Thoat")

		choice := utils.GetPositiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			student.MenuStudent()
		case 2:
			teacher.MenuTeacher()
		case 3:
			return
		default:
			fmt.Println("❌ Lua chon khong hop le")
		}

		utils.ReadInput("\n Nhan phim Enter de tiep tuc...")
	}
}
