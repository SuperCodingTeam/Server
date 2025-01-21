package services

import (
	"fmt"
	"log"

	models "github.com/SuperCodingTeam/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBook(bookName string, author string, imagePath string) *gorm.DB {
	book := models.Book{BookName: bookName, Author: author, ImagePath: imagePath}

	result := db.Create(&book)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}

	return result
}

func ReadBook() {
	var books []models.Book
	result := db.Find(&books)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}

func UpdateBookByUUID(bookUUID uuid.UUID, book models.Book) {
	var updateBook models.Book
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
	var book models.Book
	result := db.Delete(&book)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}
