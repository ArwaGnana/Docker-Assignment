package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Serve static files (for Gandalf image)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/gandalf", gandalfHandler)
	http.HandleFunc("/colombo", colomboHandler)

	port := "8080"
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  - http://localhost:8080/gandalf")
	fmt.Println("  - http://localhost:8080/colombo")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Docker Demo App</h1>")
	fmt.Fprintf(w, "<p><a href='/gandalf'>View Gandalf</a></p>")
	fmt.Fprintf(w, "<p><a href='/colombo'>View Colombo Time</a></p>")
}

func gandalfHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Gandalf the Grey</h1>")
	fmt.Fprintf(w, "<img src='/static/gandalf.jpg' width='500'>")
}

func colomboHandler(w http.ResponseWriter, r *http.Request) {
	// Load Colombo timezone
	location, err := time.LoadLocation("Asia/Colombo")
	if err != nil {
		http.Error(w, "Error loading timezone", http.StatusInternalServerError)
		return
	}

	// Get current time in Colombo
	colomboTime := time.Now().In(location)

	fmt.Fprintf(w, "<h1>Current Time in Colombo, Sri Lanka</h1>")
	fmt.Fprintf(w, "<h2>Time: %s</h2>", colomboTime.Format("15:04:05"))
	fmt.Fprintf(w, "<h3>Date: %s</h3>", colomboTime.Format("Monday, January 2, 2006"))
	fmt.Fprintf(w, "<p>Timezone: Asia/Colombo (UTC+5:30)</p>")
}
