package library

import (
	"fmt"
	"strings"
	"time"

	"trungem.com/hoc-golang/models"
)

type Library struct {
	books        map[string]models.Book
	borrowers    map[string]models.Borrower
	transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		books:        make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.books[id]; exists {
		return fmt.Errorf("sách với Id %s đã tồn tại", id)
	}

	lib.books[id] = models.Book{
		Id:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}

	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.books))
	for _, book := range lib.books {
		books = append(books, book)
	}
	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {
	if _, exists := lib.borrowers[id]; exists {
		return fmt.Errorf("người mượn với Id %s đã tồn tại", id)
	}

	lib.borrowers[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

func (lib *Library) ListBorrowerStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrowers))

	for _, borrower := range lib.borrowers {
		borrowers = append(borrowers, borrower)
	}

	return borrowers
}

func (lib *Library) BorrowerBookStore(id, bookId, borrowerId string) error {
	if _, exists := lib.transactions[id]; exists {
		return fmt.Errorf("mượn sách với Id %s đã tồn tại", id)
	}

	if _, exists := lib.books[bookId]; !exists {
		return fmt.Errorf("😞 id %s sách không tồn tại.\n", bookId)
	}

	if _, exists := lib.borrowers[borrowerId]; !exists {
		return fmt.Errorf("😞 id %s người mượn không tồn tại.\n", borrowerId)
	}

	if lib.books[bookId].IsBorrowed {
		return fmt.Errorf("😞 sách với Id %s đã được mượn, vui lòng chọn sách khác.\n", bookId)
	}

	lib.transactions[id] = models.Transaction{
		TransactionId: id,
		BookId:        bookId,
		BorrowerId:    borrowerId,
		BorrowerDate:  time.Now(),
	}

	lib.books[bookId] = models.Book{
		Id:         bookId,
		Title:      lib.books[bookId].Title,
		Author:     lib.books[bookId].Author,
		IsBorrowed: true,
	}

	return nil
}

func (lib *Library) ListBorrowerHistoryStore(borrowerId string) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, 0, len(lib.transactions))

	if _, exists := lib.borrowers[borrowerId]; !exists {
		return transactions, fmt.Errorf("😞 id %s người mượn không tồn tại.\n", borrowerId)
	}

	for _, transaction := range lib.transactions {
		if transaction.BorrowerId == borrowerId {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

func (lib *Library) ReturnBookStore(transactionId string) error {
	if _, exists := lib.transactions[transactionId]; !exists {
		return fmt.Errorf("😞 không có mã giao dịch %s\n", transactionId)
	}

	if !lib.transactions[transactionId].ReturnDate.IsZero() {
		return fmt.Errorf("😞 Sách với mã giao dịch %s đã được trả ngày %s\n", transactionId, lib.transactions[transactionId].ReturnDate.Format("2006-01-02"))
	}

	lib.transactions[transactionId] = models.Transaction{
		TransactionId: transactionId,
		BookId:        lib.transactions[transactionId].BookId,
		BorrowerId:    lib.transactions[transactionId].BorrowerId,
		BorrowerDate:  lib.transactions[transactionId].BorrowerDate,
		ReturnDate:    time.Now(),
	}

	bookId := lib.transactions[transactionId].BookId
	lib.books[bookId] = models.Book{
		Id:         bookId,
		Title:      lib.books[bookId].Title,
		Author:     lib.books[bookId].Author,
		IsBorrowed: false,
	}

	return nil
}

func (lib *Library) SearchBookStore(search string) []models.Book {
	var books []models.Book

	query := strings.ToLower(search)

	for _, book := range lib.books {
		if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Author), query) {
			books = append(books, book)
		}
	}

	return books
}
