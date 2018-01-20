package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT             = ":8000"
	STATIC_PATH      = "/static/"
	STATIC_FOR_MEDIA = "/static"
	DIR              = "./dist/static"
	MEDIA_FOLDER     = "/media/"
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

func ReturnErrorFromHandler(w http.ResponseWriter) {
	bytes, err := json.Marshal(&map[string]interface{}{
		"success": false,
	})
	if err != nil {
		log.Println("Error parse error map", err)
	}
	w.Write(bytes)
}
