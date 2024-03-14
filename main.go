package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    fmt.Println("bonjoor")
    PORT := ":3002"
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

    r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name != "" {
            w.Write([]byte("Hello " + name))
        } else {
            w.Write([]byte("Hello"))
        }
    })

    http.ListenAndServe(PORT, r)
}
