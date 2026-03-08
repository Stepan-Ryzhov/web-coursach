package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, nil)
	})

	fmt.Println("Сервер запущен на порту " + port)
	http.ListenAndServe(":"+port, nil)
}
