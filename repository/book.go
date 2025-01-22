package repository

import (
	"fmt"
	"log"

	"github.com/SuperCodingTeam/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBook(bookName string, author string, imagePath string) *gorm.DB {
	book := model.Book{BookName: bookName, Author: author, ImagePath: imagePath}

	result := db.Create(&book)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}

	return result
}

func ReadBook() {
	var books []model.Book
	result := db.Find(&books)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}

func UpdateBookByUUID(bookUUID uuid.UUID, book model.Book) {
	var updateBook model.Book
	result := db.Where("book_uuid = ?", bookUUID).First(&updateBook)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	} else {
		updateBook.BookName = book.BookName
		updateBook.Author = book.Author
		updateBook.ImagePath = book.ImagePath
	}

	db.Save(&book)
}

func DeleteBookByUUID(bookUUID uuid.UUID) {
	var book model.Book
	result := db.Delete(&book)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}
