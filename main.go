package main

import (
    "fmt"
    "net/http"
    "time"
    "ApiRateLimiterApp/Middleware"

    "github.com/gorilla/mux"
)

func main() {
    // Define rate limit parameters
    maxRequestsPerMinute := 5
    refillRate := time.Minute / time.Duration(maxRequestsPerMinute)

    // Initialize router
    r := mux.NewRouter()

    // Apply the rate limiting middleware
    r.Use(middleware.RateLimitMiddleware(maxRequestsPerMinute, refillRate))

    // Example API endpoint
    r.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "API request successful!")
    }).Methods("GET")

	// Serve static files (HTML, CSS, JS)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

    // Start the server
    http.ListenAndServe(":8080", r)
}
