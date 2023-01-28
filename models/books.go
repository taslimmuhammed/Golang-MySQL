package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model 
	ID int `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(30)" json:"firstname" binding:"required" `
	Author string `gorm:"type:varchar(30)`
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