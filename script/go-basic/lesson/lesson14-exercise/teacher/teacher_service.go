package teacher

import (
	"fmt"

	"trungem.com/hoc-golang/utils"
)

var teacherList []Teacher

func addTeacher() {
	var id int
	fmt.Println("-=-=-=-= Them giang vien -=-=-=-=")
	for {
		id = utils.GetPositiveInt("- Nhap id: ")
		if utils.IsIdUnique(id, teacherList) {
			break
		}

		fmt.Println("❌ Id da ton tai, vui long nhap id khac.")
	}

	name := utils.GetNonEmptyString("- Nhap ten: ")
	subject := utils.GetNonEmptyString("- Nhap mon giang day: ")
	baseSalary := utils.GetPositiveFloat("- Nhap luong co ban: ")
	bonus := utils.GetPositiveFloat("- Nhap thuong: ")

	teacher := Teacher{
		Id:         id,
		Name:       name,
		Subject:    subject,
		BaseSalary: baseSalary,
		Bonus:      bonus,
	}
	teacherList = append(teacherList, teacher)

	fmt.Printf("%+v \n", teacher)

	fmt.Println("✅ Them giang vien thanh cong!")
}

func deleteTeacher() {
	fmt.Println("-=-=-=-= Xoa giang vien -=-=-=-=")
	id := utils.GetPositiveInt("🗑️ Nhap ID giang vien can xoa: ")

	for idx, student := range teacherList {
		if student.Id == id {
			teacherList = append(teacherList[:idx], teacherList[idx+1:]...)
			fmt.Println("✅ Xoa giang vien thanh cong")
			return
		}
	}

	fmt.Println("❌ Khong tim thay giang vien")
}

func updateTeacher() {
	fmt.Println("-=-=-=-= Sua giang vien -=-=-=-=")
	id := utils.GetPositiveInt("✏️ Nhap ID giang vien can sua: ")

	for _, t := range teacherList {
		if t.Id == id {
			fmt.Println("🔁 Nhap thong tin moi (Nhan Enter de giu nguyen gia tri hien tai)")
			name := utils.GetOptionalString(fmt.Sprintf("- Nhap ten (%s): ", t.Name), t.Name)
			subject := utils.GetOptionalString(fmt.Sprintf("- Nhap lop (%s): ", t.Subject), t.Subject)
			baseSalary := utils.GetOptionalPositiveFloat(fmt.Sprintf("- Nhap lop (%s): ", t.Subject), t.BaseSalary)
			bonus := utils.GetOptionalPositiveFloat(fmt.Sprintf("- Nhap lop (%s): ", t.Subject), t.Bonus)

			teacher := Teacher{
				Id:         id,
				Name:       name,
				Subject:    subject,
				BaseSalary: baseSalary,
				Bonus:      bonus,
			}

			fmt.Printf("%+v \n", teacher)

			fmt.Println("✅ Cap nhat giang vien thanh cong")

			return
		}
	}

	fmt.Println("❌ Khong tim thay giang vien")
}

func listTeacher() {
	fmt.Println("-=-=-=-= Danh sach giang vien -=-=-=-=")
	if len(teacherList) == 0 {
		fmt.Println("Khong co giang vien nao trong danh sach")
		return
	}

	for _, teacher := range teacherList {
		fmt.Println(teacher.GetInfo())
	}
}

func searchTeacher() {
	fmt.Println("-=-=-=-= Tim kiem giang vien -=-=-=-=")
	id := utils.GetPositiveInt("🔎 Nhap ID giang vien can tim: ")

	for _, teacher := range teacherList {
		if teacher.Id == id {
			fmt.Println("✅ Tim thay giang vien: ", teacher.GetInfo())
			return
		}
	}

	fmt.Println("❌ Khong tim thay giang vien")
}

func MenuTeacher() {
	for {
		utils.ClearScreen()

		fmt.Println("==== QUAN LY GIANG VIEN ====")
		fmt.Println("1. Them giang vien")
		fmt.Println("2. Xoa giang vien")
		fmt.Println("3. Sua giang vien")
		fmt.Println("4. Danh sach giang vien")
		fmt.Println("5. Tim kiem giang vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPositiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			addTeacher()
		case 2:
			deleteTeacher()
		case 3:
			updateTeacher()
		case 4:
			listTeacher()
		case 5:
			searchTeacher()
		case 6:
			return
		default:
			fmt.Println("❌ Lua chon khong hop le")
		}

		utils.ReadInput("\n Nhan phim Enter de tiep tuc...")
	}
}
