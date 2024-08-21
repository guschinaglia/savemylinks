package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gxtech/savemylink/main/models"
)

type Link struct {
	gorm.Model `json:"model"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	Category   Category
}

type Category struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
}

func main() {
	gin.ForceConsoleColor()

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to Database")
	}

	migrateDb(db)
	handleRequest(db)
}

func handleRequest(db *gorm.DB) {
	r := gin.Default()
	r.GET("/users", models.AllUsers(db))
	r.POST("/users", models.CreateUser(db))
	r.Run()
}

func migrateDb(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &Link{}, &Category{})
}

//func allUsers(db *gorm.DB) func(c gin.Context) {
//	return func(c gin.Context) {
//		var users []User
//		db.Find(&users)
//		fmt.Println("{}", users)
//
//		return c.JSON(http.StatusOK, users)
//	}
//}
