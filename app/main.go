package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "database/sql"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}
var db *gorm.DB
var err error
func main() {
	db, err = gorm.Open("mysql", "docker:docker@tcp(db:3306)/test_db?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	createTable()
	r := gin.Default()
	users := r.Group("/users")
	users.GET("/:id", getUser)
	users.POST("/edit", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "user.Name", "age": "user.Age"})
	})
	r.Run()
}

func createTable() {
	db.Exec("TRUNCATE TABLE users;")
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "phule", Age: 25})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	db.First(&user, id)
	log.Println(user)
	log.Println(user.Name)
	log.Println(user.Age)
	fmt.Println("asaa1")
	c.JSON(200, gin.H{"name": user.Name, "age": user.Age})
}