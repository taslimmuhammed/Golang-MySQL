package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/taslimmuhammed/gorm/database"
	"github.com/taslimmuhammed/gorm/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type BookRepo struct{
	Db *gorm.DB
}

func NewBookRepo() *BookRepo{
  db := database.InitDb()
  db.AutoMigrate(&models.Book{})
  fmt.Println(db)
  return &BookRepo{Db: db}
}

func(repository *BookRepo) CreateBook(c *gin.Context){
   var book models.Book
    c.BindJSON(&book)
//    if er !=nil{
// 	c.AbortWithStatusJSON(500, gin.H{"error":er})
// 	return
//    }
   err := models.CreateBook(repository.Db, &book)
   if err !=nil{
	  c.AbortWithStatusJSON(500, gin.H{"error":err})
	  return
   }
   c.JSON(http.StatusOK, book)
}

func(repository *BookRepo) GetBooks(c *gin.Context ){
	var book []models.Book
	err := models.GetBooks(repository.Db, &book)
	if err !=nil{
	   c.AbortWithStatusJSON(500, gin.H{"error":err})
	   return
	}
	c.JSON(http.StatusOK, book)
}
func(repository *BookRepo) GetBook(c *gin.Context ){
	id,_ := strconv.Atoi(c.Param("id"))
    var book models.Book
   err := models.GetBook(repository.Db, &book, id)
   if err !=nil{
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	  c.AbortWithStatusJSON(500, gin.H{"error":err})
	  return
   }
   c.JSON(http.StatusOK, book)
}
func(repository *BookRepo) UpdateBook(c *gin.Context ){
	var book models.Book
	err := models.UpdateBook(repository.Db, &book)
	if err !=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	   c.AbortWithStatusJSON(500, gin.H{"error":err})
	   return
	}
	c.JSON(http.StatusOK, book)
}

func(repository *BookRepo) DeleteBook(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	var book models.Book
	err := models.DeleteBook(repository.Db,&book, id )
	if err !=nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	   c.AbortWithStatusJSON(500, gin.H{"error":err})
	   return
	}
	c.JSON(http.StatusOK, book)
} 