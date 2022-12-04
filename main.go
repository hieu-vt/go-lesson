package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

type NoteUpdate struct {
	Title *string `json:"title" gorm:"column:title;"`
}

func (Note) TableName() string {
	return "notes"
}

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if error := runService(db); error != nil {
		log.Fatalln(error)
	}

	// inset data to database
	//newNote := Note{Title: "This is domo note", Content: "This is content of demo note"}
	//
	//if err := db.Create(&newNote); err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(newNote)
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
