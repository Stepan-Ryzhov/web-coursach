package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Привет! Сервер работает")
    })

    fmt.Println("Сервер запущен на порту " + port)
    http.ListenAndServe(":"+port, nil)
}