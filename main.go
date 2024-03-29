package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"website/handlers"
	"website/middleware"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for /")
		http.ServeFile(w, r, "frontend/index.html")
	})

	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST", "GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST", "GET")
	r.HandleFunc("/profile", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/delete", handlers.DeleteHandler).Methods("POST")
	r.HandleFunc("/edit", handlers.EditHandler).Methods("GET")
	r.HandleFunc("/update", handlers.UpdateHandler).Methods("POST")
	r.HandleFunc("/filtered-books", handlers.FilterBooksHandler).Methods("GET")
	r.HandleFunc("/confirm", handlers.ConfirmHandler).Methods("GET", "POST")
	r.HandleFunc("/email-confirmed", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/email-confirmed.html")
	})
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/message", handlers.MessageHandler).Methods("GET", "POST")
	r.HandleFunc("/admin", handlers.AdminDashboardHandler).Methods("GET")
	r.HandleFunc("/edit-book", handlers.EditBookHandler).Methods("GET")
	r.HandleFunc("/delete-book", handlers.DeleteBookHandler).Methods("POST")
	r.HandleFunc("/add-book", handlers.AddBookHandler).Methods("POST")
	r.Use(middleware.RateLimitMiddleware)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
