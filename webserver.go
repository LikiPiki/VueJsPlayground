package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT        = ":8000"
	STATIC_PATH = "/static/"
	DIR         = "./dist/static"
)

func startWebRestApi() {
	router := mux.NewRouter()

	router.PathPrefix(STATIC_PATH).Handler(http.StripPrefix(STATIC_PATH, http.FileServer(http.Dir(DIR))))

	router.HandleFunc("/add_post", addNewPost).Methods("Post")
	router.HandleFunc("/login", loginHandler).Methods("Post")
	router.HandleFunc("/register", registerHandler).Methods("Post")

	router.HandleFunc("/get_posts", getPosts).Methods("Get")
	router.HandleFunc("/", homeHandler).Methods("Get")
	http.Handle("/", router)
	fmt.Println(
		fmt.Sprintf("Running server on port http://localhost%s", PORT),
	)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic("Error start web server")
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	var posts, newPosts []Post
	// var post Post

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
	if err != nil {
		log.Println(err)
		return
	} else {
		var user User
		if db.Where("username = ?", data.Username).First(&user).RecordNotFound() {
			log.Println(err)
			return
		} else {
			fmt.Println("HERE")

			db.Create(&Post{
				User:      user,
				Title:     data.Title,
				Content:   data.Content,
				ImageLink: data.ImageLink,
			})
		}

		resp, err := json.Marshal(&map[string]interface{}{
			"success": true,
		})

		if err != nil {
			log.Println(err)
		}
		w.Write(resp)
	}
	fmt.Println(data)
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
			"userImage": "",
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
		bytes, err = json.Marshal(&map[string]interface{}{
			"success": false,
		})

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
