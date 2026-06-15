package main

import (
	"fmt"

	"trungem.com/hoc-golang/library"
	"trungem.com/hoc-golang/utils"
)

func main() {
	lib := library.NewLibrary()

	for {
		utils.ClearScreen()

		fmt.Println("📚 CHƯƠNG TRÌNH QUẢN LÝ THƯ VIỆN")
		fmt.Println("──────────────────────────────────────────")
		fmt.Println("1️⃣  ➕ Thêm sách 📖")
		fmt.Println("2️⃣  📋 Xem danh sách sách 🗂️")
		fmt.Println("3️⃣  🙋‍♂️ Thêm người mượn 👩‍🏫")
		fmt.Println("4️⃣  👀 Xem danh sách người mượn 🧾")
		fmt.Println("5️⃣  📦 Mượn sách 📚")
		fmt.Println("6️⃣  🕓 Xem lịch sử mượn 📜")
		fmt.Println("7️⃣  🔁 Trả sách ✅")
		fmt.Println("8️⃣  🔍 Tìm kiếm sách 🔎")
		fmt.Println("9️⃣  🚪 Thoát chương trình 👋")
		fmt.Println("──────────────────────────────────────────")

		choice := utils.GetPositiveInt("👉 Chon chuc nang: ")

		utils.ClearScreen()

		switch choice {
		case 1:
			fmt.Println("-=-=-=-=-=-➕ Thêm sách 📖-=-=-=-=-=-")
			if err := library.AddBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi thêm sách: %v\n", err)
			}
		case 2:
			fmt.Println("-=-=-=-=-=-📋 Xem danh sách sách 🗂️-=-=-=-=-=-")
			if err := library.ListBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi xuất danh sách sách: %v\n", err)
			}
		case 3:
			fmt.Println("-=-=-=-=-=-🙋‍♂️ Thêm người mượn 👩‍🏫-=-=-=-=-=-")
			if err := library.AddBorrower(lib); err != nil {
				fmt.Printf("❌ Lỗi khi thêm người mượn: %v\n", err)
			}
		case 4:
			fmt.Println("-=-=-=-=-=-👀 Xem danh sách người mượn 🧾-=-=-=-=-=-")
			if err := library.ListBorrower(lib); err != nil {
				fmt.Printf("❌ Lỗi khi xuất danh sách người mượn: %v\n", err)
			}
		case 5:
			fmt.Println("-=-=-=-=-=-📦 Mượn sách 📚-=-=-=-=-=-")
			if err := library.BorrowerBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi thêm mượn sách: %v\n", err)
			}
		case 6:
			fmt.Println("-=-=-=-=-=-🕓 Xem lịch sử mượn 📜-=-=-=-=-=-")
			if err := library.ListBorrowerHistory(lib); err != nil {
				fmt.Printf("❌ Lỗi khi xuất danh sách lịch sử mượn: %v\n", err)
			}
		case 7:
			fmt.Println("-=-=-=-=-=-🔁 Trả sách ✅-=-=-=-=-=-")
			if err := library.ReturnBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi trả sách: %v\n", err)
			}
		case 8:
			fmt.Println("-=-=-=-=-=-🔍 Tìm kiếm sách 🔎-=-=-=-=-=-")
			if err := library.SearchBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi tìm kiếm sách: %v\n", err)
			}
		case 9:
			fmt.Println("🚀 Thoát chương trình thành công! 👋")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ.")
		}

		utils.ReadInput("\n Nhấn phím Enter để tiếp tục...")
	}
}
