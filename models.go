package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"-"`
	IsAdmin   bool   `json:"isAdmin"`
	ImagePath string `json:"imagePath"`
}

type Post struct {
	gorm.Model
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ImageLink   string    `json:"imageLink"`
	PublishTime time.Time `json:"publishTime"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user"`
}
