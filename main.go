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

type User struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

// `Profile` belongs to `User`, `UserID` is the foreign key

type Profile struct {
	Id   int    `json:"id,omitempty" gorm:"column:id"`
	User User   `gorm:"association_foreignKey:name"`
	Name string `json:"name" gorm:"column:name;"`
}

func main() {
	os.Setenv("DBConnectionStr", "root:ead8686ba57479778a76e@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// insert profile of user
	newProfile := Profile{Name: "hieu thu hai", User: User{Name: "Hieu Vu"}}

	if errCrProfile := db.Create(&newProfile); errCrProfile != nil {
		log.Fatalln(newProfile)
	}

	fmt.Println(newProfile)

	// insert data user to database
	//newUser := User{
	//	Name: "Vu Thanh Hieu",
	//}
	//
	//if errCrUser := db.Create(&newUser); errCrUser != nil {
	//	log.Fatalln(errCrUser)
	//}
	//
	//fmt.Println(newUser)

	// inset data to database
	//newNote := Note{Title: "This is domo note", Content: "This is content of demo note"}
	//
	//if err := db.Create(&newNote); err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(newNote)

	// get data
	//var notes []Note
	//
	//db.Where("status = ?", 1).Find(&notes)
	//
	//fmt.Println("Notes: ", notes)
	//
	//var note Note
	//
	//result := db.Where("id = ?", 3).First(&note)
	//
	//if err := result.Error; err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println("Result: ", note)

	// delete item
	//db.Table(Note{}.TableName()).Where("id = ?", 1).Delete(nil)

	// update
	//db.Table(Note{}.TableName()).Where("id = 2").Updates(map[string]interface{}{
	//	"title": "Demo 1",
	//})
	//note.Title = "Demo 2"
	//db.Table(Note{}.TableName()).Where("id = 3").Updates(&note)
	//
	//newTitle := "Demo 4   "
	//db.Table(Note{}.TableName()).Where("id = 3").Updates(&NoteUpdate{Title: &newTitle})
}
