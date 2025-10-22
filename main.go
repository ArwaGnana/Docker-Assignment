package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/home.html")
}

func greetUser(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		user = "Guest"
	}
	message := fmt.Sprintf("Welcome, %s!", user)
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/greet", greetUser)

	fmt.Println("Server listening on http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
