package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

// imageLoader
type ImageLoad struct {
	ImageData string `json:"imageData"`
	ImageName string `json:"imageName"`
}

func (i ImageLoad) saveImage() (err error) {
	content := strings.Split(i.ImageData, ",")
	var data []byte

	data, err = base64.StdEncoding.DecodeString(content[1])
	if err != nil {
		log.Println("Erorr decode base64 image")
		return
	}
	if i.ImageData != "" && i.ImageName != "" {
		var f *os.File
		f, err = os.Create(
			fmt.Sprintf("%s%s%s", DIR, MEDIA_FOLDER, i.ImageName),
		)
		if err != nil {
			return
		}
		f.Close()
		f.Write(data)
	}
	return
}
