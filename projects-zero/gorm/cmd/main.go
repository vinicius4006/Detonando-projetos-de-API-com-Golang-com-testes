package main

import (
	"fmt"
	"log"

	model "example.com/gorm-project/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	databaseName := "postgres"
	userName := "postgres"
	password := "123456789"
	host := "localhost"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable timeZone=America/Sao_Paulo",
		host, userName, password, databaseName)

	fmt.Println(dsn)

	//Open connection to a postgresql database:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Panic("failedto connect database")
	}
	fmt.Println("OK")

	// Let's run the migrations to create tables:

	db.Migrator().AutoMigrate(&model.Member{}, &model.Note{})

	// Let's create the constrants
	db.Migrator().CreateConstraint(&model.Member{}, "Notes")
	db.Migrator().CreateConstraint(&model.Member{}, "Connections")

	// Let's add some data:
	user1 := &model.Member{Name: "Paul", Email: "paul@testmail"}
	db.Create(&user1)
	user2 := &model.Member{Name: "John", Email: "john@test"}
	db.Create(&user2)

	// Connect John to Paul
	db.Model(&user2).Association("Connections").Append(user1)
	db.Save(user2)

	// Paul wrote a note:
	db.Model(&user1).Association("Notes").Append(&model.Note{Author: *user1, Text: "This is a note"})
	db.Save(user1)

	paul := &model.Member{}
	db.Preload("Notes").Preload("Connections").First(&paul, 1)
	fmt.Printf("\nPaul has a note with text: %s\n", paul.Notes[0].Text)
	fmt.Printf("\nPaul has %d connections\n", len(paul.Connections))

	// Now get user John, whois connected to Paul:
	john := &model.Member{}
	db.Preload("Notes").Preload("Connections").First(&john, 2)
	fmt.Printf("\nJohn has %d notes\n", len(john.Notes))
	fmt.Printf("\nJohn has a connection with %s\n", john.Connections[0].Name)

}
