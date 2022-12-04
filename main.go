package main

import (
	"fmt"
	"log"
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

	// inset data to database
	//newNote := Note{Title: "This is domo note", Content: "This is content of demo note"}
	//
	//if err := db.Create(&newNote); err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(newNote)

	// get data
	var notes []Note

	db.Where("status = ?", 1).Find(&notes)

	fmt.Println("Notes: ", notes)

	var note Note

	result := db.Where("id = ?", 3).First(&note)

	if err := result.Error; err != nil {
		log.Println(err)
	}

	fmt.Println("Result: ", note)

	// delete item
	//db.Table(Note{}.TableName()).Where("id = ?", 1).Delete(nil)

	// update
	//db.Table(Note{}.TableName()).Where("id = 2").Updates(map[string]interface{}{
	//	"title": "Demo 1",
	//})
	//note.Title = "Demo 2"
	//db.Table(Note{}.TableName()).Where("id = 3").Updates(&note)

	newTitle := "Demo 4   "
	db.Table(Note{}.TableName()).Where("id = 3").Updates(&NoteUpdate{Title: &newTitle})
}
