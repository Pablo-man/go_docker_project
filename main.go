package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola, mundo desde Go!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Servidor ejecutándose en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
