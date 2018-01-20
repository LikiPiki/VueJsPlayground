package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"-"`
	IsAdmin   bool   `json:"isAdmin"`
	ImagePath string `json:"imagePath"`
}

func (u User) getUserByUsername(username string) (err error) {
	if db.Where("username = ?", username).First(&u).RecordNotFound() {
		log.Println("error find user")
		return errors.New("Erorr find user from db")
	}
	return
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

func (p Post) getAllPosts() (posts []Post) {
  // add return all posts here
  return
}

// imageLoader
type ImageLoad struct {
	ImageData string `json:"imageData"`
	ImageName string `json:"imageName"`
}

func (i ImageLoad) saveImage(username string) (err error) {
	content := strings.Split(i.ImageData, ",")[1]
	var data []byte

	data, err = base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Println("Erorr decode base64 image", err)
		return
	}

	if i.ImageData != "" && i.ImageName != "" {
		var f *os.File
		f, err = os.Create(
			fmt.Sprintf("%s%s%s-%s", DIR, MEDIA_FOLDER, username, i.ImageName),
		)
		if err != nil {
			log.Println("Error", err)
			return
		}
		f.Write(data)
		f.Close()
	}
	return
}
