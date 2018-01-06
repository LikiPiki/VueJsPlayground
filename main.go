package main

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("Error open data.db")
	}

	db.AutoMigrate(
		&User{},
	)

	db.LogMode(true)

}

func main() {
	defer db.Close()

	var users []User
	db.Find(&users)
	for _, el := range users {
		fmt.Println("%#v", el)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go startWebRestApi()

	wg.Wait()

}
