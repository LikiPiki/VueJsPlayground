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
	router.HandleFunc("/login", loginHandler).Methods("Post")
	router.HandleFunc("/register", registerHandler).Methods("Post")
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
			"success":  true,
			"username": user.Username,
			"isAdmin":  user.IsAdmin,
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
