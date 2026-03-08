package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type RepairRequest struct {
	Car      string `json:"car"`
	Problem  string `json:"problem"`
	Priority string `json:"priority"`
	Mechanic string `json:"mechanic"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/emb.ico")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/home.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Главная " + time.Now().String())
	})

	http.HandleFunc("/cars.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/cars.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Автомобили " + time.Now().String())
	})

	http.HandleFunc("/drivers.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/drivers.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Водители " + time.Now().String())
	})

	http.HandleFunc("/mans.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/mans.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Механики " + time.Now().String())
	})

	http.HandleFunc("/remont.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/remont.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Ремонт " + time.Now().String())
	})

	http.HandleFunc("/routes.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/routes.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Маршруты " + time.Now().String())
	})

	http.HandleFunc("/zapch.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/zapch.html")
		tmpl.Execute(w, nil)
		fmt.Println("Обращение к странице Запчасти " + time.Now().String())
	})

	http.HandleFunc("/api/repair", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req RepairRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		fmt.Println("Новая заявка на ремонт:")
		fmt.Println("Автомобиль:  " + req.Car)
		fmt.Println("Проблема:    " + req.Problem)
		fmt.Println("Приоритет:   " + req.Priority)
		fmt.Println("Механик:     " + req.Mechanic)
		fmt.Println("Время:       " + time.Now().String())

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	fmt.Println("Сервер запущен на порту " + port)
	http.ListenAndServe(":"+port, nil)
}
