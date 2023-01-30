package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model 
	ID int `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(30)" json:"firstname" binding:"required" `
	Author string `gorm:"type:varchar(30)"`
}

func CreateBook(db *gorm.DB, Book *Book) (err error) {
	err = db.Create(Book).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBooks(db *gorm.DB, book *[]Book)(err error){

		fmt.Println("Books is nil" ,book)
		fmt.Println("the db is ", db)
    err = db.Find(book).Error
	fmt.Println(book)
	if err != nil{
		return err
	}
	return nil
}
func GetBook(db *gorm.DB, book *Book, id int)(err error){
    err = db.Where("id = ?", id).First(book).Error
	if err != nil{
		return err
	}
	return nil
}
func UpdateBook(db *gorm.DB, book *Book)(err error){
    err = db.Save(book).Error
	if err != nil{
		return err
	}
	return nil
}

func DeleteBook(db *gorm.DB, book *Book, id int)(err error){
    err = db.Where("id = ?", id).Delete(book).Error
	if err != nil{
		return err
	}
	return nil
}

func QueryFunction(db *gorm.DB, books *[] Book, id int) (err error){
	//get array of elements by id
   //err = db.Find(&books, []int{id,id+1}).Error
// Get the first record ordered by primary key
 //db.First(&user)

// Get one record, no specified order
//db.Take(&user)

// Get last record, ordered by primary key desc
//db.Last(&user)

//result := db.First(&user)
//result.RowsAffected // returns count of records found
//result.Error        // returns error or nil

// check error ErrRecordNotFound
//errors.Is(result.Error, gorm.ErrRecordNotFound)
//db.Where("author <> ?", "Dance").Find(&books)
//db.Where("name = ?", "Dance of dragons").First(&books)
// Struct
// db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)

   return err
}