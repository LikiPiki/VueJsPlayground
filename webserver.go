package main

import (
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("dist/index.html")
	if err != nil {
		log.Println("Error parse file index.html")
	}
	temp.ExecuteTemplate(w, "index.html", nil)
}
