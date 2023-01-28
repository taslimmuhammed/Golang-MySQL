package main

import (
   "github.com/gin-gonic/gin"
   "github.com/taslimmuhammed/gorm/controllers"
   "net/http"
)
var (
	userRepo   = controllers.New()
)
func main() {
   r := setupRouter()
   r.POST("/users", userRepo.CreateUser)
   r.GET("/users", userRepo.GetUsers)
   r.GET("/users/:id", userRepo.GetUser)
   r.PUT("/users/:id", userRepo.UpdateUser)
   r.DELETE("/users/:id", userRepo.DeleteUser)

   _ = r.Run(":8080")
   
}

func setupRouter() *gin.Engine {
   r := gin.Default()

   r.GET("ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, "pong")
   })

   return r
}