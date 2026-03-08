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

	http.HandleFunc("/cars.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/cars.html")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/drivers.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/drivers.html")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/mans.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/mans.html")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/remont.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/remont.html")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/routes.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/routes.html")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/zapch.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/zapch.html")
		tmpl.Execute(w, nil)
	})

	fmt.Println("Сервер запущен на порту " + port)
	http.ListenAndServe(":"+port, nil)
}
