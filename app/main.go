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

func main() {
	createTable()
	r := gin.Default()
	r.GET("/ping/:id", func(c *gin.Context) {
		db, err := gorm.Open("mysql", "docker:docker@tcp(db:3306)/test_db?parseTime=true")
		if err != nil {
				log.Fatal(err)
		}
		id := c.Param("id")
		var user User
		  db.First(&user, id)
          log.Println(user)
          log.Println(user.Name)
          log.Println(user.Age)
            fmt.Println("asaa")
		c.JSON(200, gin.H{"name": user.Name, "age": user.Age})
	})

	r.Run()
}

func createTable() {
	db, err := gorm.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
	if err != nil {
	  panic("failed to connect database")
	}
    db.Exec("TRUNCATE TABLE users;")
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "phule", Age: 25})
}
