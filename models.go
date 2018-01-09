package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type User struct {
	gorm.Model
	Username  string
	Password  string `json:"-"`
	IsAdmin   bool
	ImagePath string
}

type Post struct {
	gorm.Model
	Title       string
	Content     string
	ImageLink   string
	PublishTime time.Time
	UserID      int
	User        User
}
