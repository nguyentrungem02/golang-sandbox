package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"trungem.com/hoc-gin/internal/db"
	"trungem.com/hoc-gin/internal/handler"
	"trungem.com/hoc-gin/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if err := db.InitDB(); err != nil {
		log.Fatalln("unable to connect to database")
	}
	log.Println(db.DB)

	r := gin.Default()

	userRepository := repository.NewSQLUserRepository()
	userHandler := handler.NewUserHandler(userRepository)

	r.GET("/api/v1/users/:id", userHandler.GetUserById)
	r.POST("/api/v1/users", userHandler.CreateUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
