package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
  "strings"
  "encoding/base64"
  "os"
)

const (
	PORT        = ":8000"
	STATIC_PATH = "/static/"
	STATIC_FOR_MEDIA ="/static"
	DIR         = "./dist/static"
	MEDIA_FOLDER = "/media/"
)

func startWebRestApi() {
	router := mux.NewRouter()

	router.PathPrefix(STATIC_PATH).Handler(http.StripPrefix(STATIC_PATH, http.FileServer(http.Dir(DIR))))

	router.HandleFunc("/get_user/{username}", getUser).Methods("Get")
	router.HandleFunc("/add_post", addNewPost).Methods("Post")
	router.HandleFunc("/load_profile_image", loadProfileImage).Methods("Post")
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

func getUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  username := vars["username"]
  var user User
  if db.Where("username = ?", username).First(&user).RecordNotFound() {
    // error here
  } else {
    bytes, err := json.MarshalIndent(&user, "", "\t")
    if err != nil {
      // error here
    }
    w.Write(bytes)
  }
}

func loadProfileImage(w http.ResponseWriter, r *http.Request) {
  //save image from query
  type Body struct {
    Username string `json:"username"`
    ImageData string `json:"imageData"`
    ImageName string `json:"imageName"`
  }
  var body Body
  err := json.NewDecoder(r.Body).Decode(&body)
  if err != nil {
    // error here
  }
  fmt.Println("Username is ", body.Username)
  content := strings.Split(body.ImageData, ",")
  data, err := base64.StdEncoding.DecodeString(content[1])
  if err != nil {
    log.Println("Erorr decode base64 image")
  }
  if body.ImageData != "" && body.ImageName != "" {
    f, err := os.Create(
      fmt.Sprintf("%s%s%s", DIR, MEDIA_FOLDER,  body.ImageName),
    )
    if err != nil {
      log.Println("error writing file")
      // erorr here
    }
    f.Write(data)
    f.Close()
  }
  var user User
  if db.Where("username = ?", body.Username).First(&user).RecordNotFound() {
    // error here
  } else {
    db.Model(&user).Update("imagePath", fmt.Sprintf("%s%s%s", STATIC_FOR_MEDIA, MEDIA_FOLDER,  body.ImageName))
  }
  // return something cool!!!
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
