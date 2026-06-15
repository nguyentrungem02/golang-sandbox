package student

import (
	"fmt"

	"trungem.com/hoc-golang/utils"
)

var studentList []Student

func addStudent() {
	var scores []float64
	var id int
	fmt.Println("-=-=-=-= Them sinh vien -=-=-=-=")
	for {
		id = utils.GetPositiveInt("- Nhap id: ")
		if utils.IsIdUnique(id, studentList) {
			break
		}

		fmt.Println("❌ Id da ton tai, vui long nhap id khac.")
	}

	name := utils.GetNonEmptyString("- Nhap ten: ")
	class := utils.GetNonEmptyString("- Nhap lop: ")
	totalPoint := utils.GetPositiveInt("- Nhap so luong diem: ")

	for i := 1; i <= totalPoint; i++ {
		score := utils.GetPositiveFloat(fmt.Sprintf("- Nhap diem %d: ", i))
		scores = append(scores, score)
	}

	student := Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	}
	studentList = append(studentList, student)

	fmt.Printf("%+v \n", student)

	fmt.Println("✅ Them sinh vien thanh cong")
}

func deleteStudent() {
	fmt.Println("-=-=-=-= Xoa sinh vien -=-=-=-=")
	id := utils.GetPositiveInt("🗑️ Nhap ID sinh vien can xoa: ")

	for idx, student := range studentList {
		if student.Id == id {
			studentList = append(studentList[:idx], studentList[idx+1:]...)
			fmt.Println("✅ Xoa sinh vien thanh cong")
			return
		}
	}

	fmt.Println("❌ Khong tim thay sinh vien")
}

func updateStudent() {
	fmt.Println("-=-=-=-= Sua sinh vien -=-=-=-=")
	id := utils.GetPositiveInt("✏️ Nhap ID sinh vien can sua: ")

	for _, s := range studentList {
		if s.Id == id {
			fmt.Println("🔁 Nhap thong tin moi (Nhan Enter de giu nguyen gia tri hien tai)")
			name := utils.GetOptionalString(fmt.Sprintf("- Nhap ten (%s): ", s.Name), s.Name)
			class := utils.GetOptionalString(fmt.Sprintf("- Nhap lop (%s): ", s.Class), s.Class)

			newScores := make([]float64, len(s.Scores))
			for i, score := range s.Scores {
				newScores[i] = utils.GetOptionalPositiveFloat(fmt.Sprintf("- Nhap diem %d (%.2f): ", i+1, score), score)
			}

			student := Student{
				Id:     id,
				Name:   name,
				Class:  class,
				Scores: newScores,
			}

			fmt.Printf("%+v \n", student)

			fmt.Println("✅ Cap nhat sinh vien thanh cong")

			return
		}
	}

	fmt.Println("❌ Khong tim thay sinh vien")
}

func listStudent() {
	fmt.Println("-=-=-=-= Danh sach sinh vien -=-=-=-=")
	if len(studentList) == 0 {
		fmt.Println("Khong co sinh vien nao trong danh sach")
		return
	}
	for _, student := range studentList {
		fmt.Println(student.GetInfo())
	}

}

func searchStudent() {
	fmt.Println("-=-=-=-= Tim kiem sinh vien -=-=-=-=")
	id := utils.GetPositiveInt("🔎 Nhap ID sinh vien can tim: ")

	for _, student := range studentList {
		if student.Id == id {
			fmt.Println("✅ Tim thay sinh vien: ", student.GetInfo())
			return
		}
	}

	fmt.Println("❌ Khong tim thay sinh vien")
}

func MenuStudent() {
	for {
		utils.ClearScreen()

		fmt.Println("==== QUAN LY SINH VIEN ====")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPositiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			updateStudent()
		case 4:
			listStudent()
		case 5:
			searchStudent()
		case 6:
			return
		default:
			fmt.Println("❌ Lua chon khong hop le")
		}

		utils.ReadInput("\n Nhan phim Enter de tiep tuc...")
	}
}
