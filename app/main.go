package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
	"log"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}
type UserTable struct {
	gorm.Model
	Name string
	Age  int64
}
var db *gorm.DB
var err error
type UserRequest struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	Age string `form:"age" json:"age" xml:"age" binding:"required"`
}

type testHeader struct {
	Authorization string `header:"Authorization"`
}

func main() {
	db, err = gorm.Open("mysql", "docker:docker@tcp(db:3306)/test_db?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	createTable()
	r := gin.Default()
	r.GET("get/:id/:ten", func(c *gin.Context) {
		id := c.Param("id")
		ten := c.Param("ten")

		c.JSON(200, gin.H{"id": id, "ten": ten})
	})
	users := r.Group("/users")
	{
		users.GET("/:id", getUser)
		users.POST("/create", createUser)
		users.PUT("/edit", editUser)
		users.DELETE("/delete", deleteUser)
	}
	r.Run(":8081")
}

func createTable() {
	db.Exec("DROP TABLE users;")
	db.Exec("DROP TABLE user_tables;")
	db.AutoMigrate(&User{}, &UserTable{})
	db.Create(&User{Name: "phule", Age: 25})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	db.First(&user, id)
	c.JSON(200, gin.H{"name": user.Name, "age": user.Age})
}

func createUser(c *gin.Context) {
	var json UserRequest
	h := testHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(400, err)
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(h.Authorization)
	log.Println(json.Age)
	c.JSON(200, gin.H{"name": json.Name, "age": json.Age})
}

func editUser(c *gin.Context) {
	var json UserRequest
	h := testHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(400, err)
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(h.Authorization)
	log.Println(json.Age)
	c.JSON(200, gin.H{"name": json.Name, "age": json.Age})
}

func deleteUser(c *gin.Context) {
	var json UserRequest
	h := testHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(400, err)
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(h.Authorization)
	log.Println(json.Age)
	c.JSON(200, gin.H{"name": json.Name, "age": json.Age})
}