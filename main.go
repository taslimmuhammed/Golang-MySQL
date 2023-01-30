package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taslimmuhammed/gorm/controllers"
)
var (
	userController  = controllers.NewUserRepo()
   bookController = controllers.NewBookRepo()
)

func main() {
   r := setupRouter()
   userRoutes := r.Group("/user")

   {
   userRoutes.POST("/", userController.CreateUser)
   userRoutes.GET("/", userController.GetUsers)
   userRoutes.GET("/:id", userController.GetUser)
   userRoutes.PUT("/:id", userController.UpdateUser)
   userRoutes.DELETE("/:id", userController.DeleteUser)
   }
   
   bookRoutes := r.Group("/book")

   {
   bookRoutes.POST("/", bookController.CreateBook)
   bookRoutes.GET("/", bookController.GetBooks)
   bookRoutes.GET("/:id", bookController.GetBook)
   bookRoutes.PUT("/:id", bookController.UpdateBook)
   bookRoutes.DELETE("/:id", bookController.DeleteBook)
   bookRoutes.GET("/query/:id", bookController.QueryFunctions)
   }

   err := r.Run(":8080")
   if err != nil{
      fmt.Println(err)
   }
   
}

func setupRouter() *gin.Engine {
   r := gin.New()

   r.GET("ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, "pong")
   })

   return r
}