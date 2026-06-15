package library

import (
	"fmt"

	"trungem.com/hoc-golang/utils"
)

func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNonEmptyString("- Nhập tiêu đề: ")
	author := utils.GetNonEmptyString("- Nhập tác giả: ")

	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	fmt.Printf("✅ Thêm sách thành công! ID: %s\n", id)

	return nil
}

func ListBook(lib *Library) error {
	books := lib.ListBooksStore()

	if len(books) == 0 {
		fmt.Println("🚀Danh sách trống! Vui lòng thêm sách để xem.")
		return nil
	}

	for _, book := range books {
		status := "Còn"
		if book.IsBorrowed {
			status = "Đã mượn"
		}

		fmt.Printf("Id: %s, Tiêu đề: %s, Tác giả: %s, Trạng thái: %s\n", book.Id, book.Title, book.Author, status)
	}

	return nil
}

func AddBorrower(lib *Library) error {
	id := utils.GenerateId()
	name := utils.GetNonEmptyString("- Nhập tên người mượn: ")
	var email string

	for {
		email = utils.GetNonEmptyString("- Nhập email: ")
		if utils.ValidEmail(email) {
			break
		}

		fmt.Println("❌ Email không hợp lệ.")
	}

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	fmt.Printf("✅ Thêm sách thành công! ID: %s\n", id)

	return nil
}

func ListBorrower(lib *Library) error {
	borrowers := lib.ListBorrowerStore()

	if len(borrowers) == 0 {
		fmt.Println("🚀Danh sách trống! Vui lòng thêm người mượn để xem.")
		return nil
	}

	for _, borrower := range borrowers {
		fmt.Printf("Id: %s, Tên: %s, Email: %s\n", borrower.Id, borrower.Name, borrower.Email)
	}

	return nil
}
func BorrowerBook(lib *Library) error {
	id := utils.GenerateId()
	bookId := utils.GetNonEmptyString("- Nhập ID sách cần mượn: ")
	borrowerId := utils.GetNonEmptyString("- Nhập ID người mượn: ")

	if err := lib.BorrowerBookStore(id, bookId, borrowerId); err != nil {
		return err
	}

	fmt.Printf("✅ Mượn sách thành công! ID: %s\n", id)

	return nil
}

func ListBorrowerHistory(lib *Library) error {
	borrowerId := utils.GetNonEmptyString("- Nhập ID người mượn: ")
	transactions, err := lib.ListBorrowerHistoryStore(borrowerId)

	if err != nil {
		return err
	}

	if len(transactions) == 0 {
		fmt.Println("🚀 Danh sách trống! Chưa có người mượn sách.")
		return nil
	}

	for _, transaction := range transactions {
		returnDate := "Chưa trả"
		if !transaction.ReturnDate.IsZero() {
			returnDate = transaction.ReturnDate.Format("2006-01-02")
		}
		fmt.Printf("Giao dịch: %s, Sách: %s, Ngày mượn: %s, Ngày trả: %s\n", transaction.TransactionId, transaction.BookId, transaction.BorrowerDate.Format("2006-01-02"), returnDate)
	}

	return nil
}

func ReturnBook(lib *Library) error {
	transactionId := utils.GetNonEmptyString("- Nhập mã giao dich: ")

	if err := lib.ReturnBookStore(transactionId); err != nil {
		return err
	}

	fmt.Println("✅ Trả sách thành công!")

	return nil
}

func SearchBook(lib *Library) error {
	search := utils.GetNonEmptyString("- Nhập tiêu đề hoặc tác giả để tìm kiếm: ")
	books := lib.SearchBookStore(search)

	if len(books) == 0 {
		fmt.Println("🚀 Danh sách trống! Không có tiêu đề hay tác giả nào như từ khoá bạn tìm kiếm.")
		return nil
	}

	for _, book := range books {
		status := "Còn"
		if book.IsBorrowed {
			status = "Đã mượn"
		}

		fmt.Printf("Id: %s, Tiêu đề: %s, Tác giả: %s, Trạng thái: %s\n", book.Id, book.Title, book.Author, status)
	}

	return nil
}
