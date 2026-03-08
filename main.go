package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var ProfilesList struct {
	name string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/home.html")
		tmpl.Execute(w, nil)
	})

	fmt.Println("Сервер запущен на порту " + port)
	http.ListenAndServe(":"+port, nil)
}
