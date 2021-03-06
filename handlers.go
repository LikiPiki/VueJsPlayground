package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	var user User
	err := user.getUserByUsername(username)
	if err != nil {
	  ReturnErrorFromHandler(w)
	  return
	} else {
		bytes, err := json.MarshalIndent(&user, "", "\t")
		if err != nil {
			ReturnErrorFromHandler(w)
			return
		}
		w.Write(bytes)
	}
}

func loadProfileImage(w http.ResponseWriter, r *http.Request) {
	//save image from query
	type Body struct {
		Username string `json:"username"`
		ImageLoad
	}
	var body Body
	var err error

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Error decode", err)
		ReturnErrorFromHandler(w)
		return
	}

	err = body.saveImage(body.Username)
	if err != nil {
		log.Println("Error save image", err)
		ReturnErrorFromHandler(w)
		return
	}

	var user User
	if err := user.getUserByUsername(body.Username); err != nil {
		log.Println(err)
		ReturnErrorFromHandler(w)
		return
	}
	db.Model(&user).Update("imagePath", fmt.Sprintf("%s%s%s-%s", STATIC_FOR_MEDIA, MEDIA_FOLDER, body.Username, body.ImageName))

	// write succes json to writter
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	var posts, newPosts []Post
	db.Find(&posts)

	for _, el := range posts {
		db.Model(&el).Related(&el.User)
		newPosts = append(newPosts, el)
	}
	data, err := json.MarshalIndent(&newPosts, "", "\t")
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

// change this func
func addNewPost(w http.ResponseWriter, r *http.Request) {

	type Body struct {
		Username  string
		IsAdmin   bool
		imageLink string
		Title     string
		Content   string
		ImageLink string
	}

	var data Body
	err := json.NewDecoder(r.Body).Decode(&data)
  var post Post
  var user User

	if err != nil {
		log.Println(err)
		return
	} else {
		if db.Where("username = ?", data.Username).First(&user).RecordNotFound() {
			log.Println(err)
			return
		} else {
			fmt.Println("HERE")
			post = Post{
				User:      user,
				Title:     data.Title,
				Content:   data.Content,
				ImageLink: data.ImageLink,
			}
			db.Create(&post)
		}

		fmt.Println("Result is ", post.ID)
		fmt.Println("post is", post)

		resp, err := json.Marshal(&post)

		if err != nil {
			log.Println(err)
		}
		w.Write(resp)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var err error
	var bytes []byte
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Error parse response body")
	}
	var user User
	if !db.Where(
		"username = ? and password = ?",
		data["username"].(string),
		data["password"].(string),
	).First(&user).RecordNotFound() {
		bytes, err = json.Marshal(&map[string]interface{}{
			"success":   true,
			"username":  user.Username,
			"isAdmin":   user.IsAdmin,
			"userImage": user.ImagePath,
		})
	} else {
		bytes, err = json.Marshal(&map[string]interface{}{
			"success": false,
		})
	}
	if err != nil {
		log.Println("Cant json return value: ", err)
	}
	w.Write(bytes)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var err error
	var bytes []byte
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Error parse response body")
	}
	var user User
	if db.Where("username = ?", data["username"].(string)).First(&user).RecordNotFound() {
		db.Create(
			&User{
				Username: data["username"].(string),
				Password: data["password"].(string),
				IsAdmin:  false,
			},
		)
		bytes, err = json.Marshal(&map[string]interface{}{
			"success": true,
		})
	} else {
		ReturnErrorFromHandler(w)
		return
	}
	if err != nil {
		log.Println("Error in marshal json", err)
	}
	w.Write(bytes)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("dist/index.html")
	if err != nil {
		log.Println("Error parse file index.html")
	}
	temp.ExecuteTemplate(w, "index.html", nil)
}
